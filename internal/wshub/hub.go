package wshub

import (
	"context"
	"fmt"
	"slices"
	"time"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

type Hub struct {
	clients map[*Client]bool
}

func New() *Hub {
	return &Hub{
		clients: make(map[*Client]bool),
	}
}

func (h *Hub) Broadcast(message any, clientIDs ...string) {
	for client := range h.clients {
		select {
		case client.send <- message:
		default:
			close(client.send)
			delete(h.clients, client)
		}
	}
}

func (h *Hub) Broadacst(message any, clientIDs ...string) {
	for client := range h.clients {
		if slices.Contains(clientIDs, client.id) {
			select {
			case client.send <- message:
			default:
				close(client.send)
				delete(h.clients, client)
			}
		}
	}
}

func (h *Hub) HandleClient(client *Client) error {
	h.register(client)
	defer h.unregister(client)

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return h.handleWrite(ctx, client)
	})
	g.Go(func() error {
		return h.handleRead(ctx, client)
	})

	return g.Wait()
}

func (h *Hub) register(client *Client) {
	h.clients[client] = true
}

func (h *Hub) unregister(client *Client) {
	if _, ok := h.clients[client]; ok {
		delete(h.clients, client)
		close(client.send)
	}
}

func (h *Hub) handleRead(ctx context.Context, client *Client) error {
	client.conn.SetReadLimit(maxMessageSize)
	if err := client.conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		return fmt.Errorf("failed to set connection read deadline: %w", err)
	}
	client.conn.SetPongHandler(func(string) error { return client.conn.SetReadDeadline(time.Now().Add(pongWait)) })

	child, cancel := context.WithCancel(ctx)
	var err error
	go func() {
		for {
			_, _, err = client.conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Error().Err(err).Msg("unexpected close of websocket")
				}
				break
			}
		}
		cancel()
	}()

	<-child.Done()
	return nil
}

func (h *Hub) handleWrite(ctx context.Context, client *Client) error {
	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil
		case message, ok := <-client.send:
			if err := client.conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				return err
			}
			if !ok {
				return client.conn.WriteMessage(websocket.CloseMessage, []byte{})
			}

			if err := client.conn.WriteJSON(message); err != nil {
				log.Warn().Err(err).Msg("failed to write message")
			}
		case <-ticker.C:
			if err := client.conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				return err
			}
			if err := client.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return err
			}
		}
	}
}
