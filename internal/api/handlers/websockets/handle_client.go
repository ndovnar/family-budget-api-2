package websockets

import (
	"errors"
	"net/http"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"

	"github.com/ndovnar/family-budget-api/internal/api/error"
	"github.com/ndovnar/family-budget-api/internal/wshub"
)

func (ws *Websockets) HandleClient(ctx *gin.Context) {
	token := ctx.Query("token")

	claims, err := ws.auth.VerifyToken(token)
	if err != nil {
		ctx.Error(error.NewHttpError(http.StatusUnauthorized))
		return
	}

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Warn().Err(err).Msg("failed to upgrade websocket connection")
		return
	}
	defer conn.Close()

	client := wshub.NewClient(claims.UserID, conn)

	if err := ws.hub.HandleClient(client); err != nil {
		if !errors.Is(err, syscall.EPIPE) && !errors.Is(err, websocket.ErrCloseSent) {
			log.Warn().Err(err).Msg("failed to handle websocket connection")
		}
	}

	if err := conn.Close(); err != nil {
		log.Warn().Err(err).Msg("failed to close websocket connection")
	}
}
