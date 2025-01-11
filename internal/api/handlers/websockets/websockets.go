package websockets

import (
	"github.com/ndovnar/family-budget-api/internal/auth"
	"github.com/ndovnar/family-budget-api/internal/wshub"
)

type Websockets struct {
	auth *auth.Auth
	hub  *wshub.Hub
}

func New(auth *auth.Auth, hub *wshub.Hub) *Websockets {
	return &Websockets{
		auth: auth,
		hub:  hub,
	}
}
