package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/ndovnar/family-budget-api/internal/auth"
	"github.com/ndovnar/family-budget-api/internal/authz"
	"github.com/ndovnar/family-budget-api/internal/store"
	"github.com/ndovnar/family-budget-api/internal/wshub"
)

type API struct {
	config Config
	auth   *auth.Auth
	authz  *authz.Authz
	wshub  *wshub.Hub
	router *gin.Engine
	store  store.Store
	server *http.Server
}

func New(cfg Config, auth *auth.Auth, authz *authz.Authz, wshub *wshub.Hub, store store.Store) *API {
	router := gin.Default()

	api := &API{
		config: cfg,
		auth:   auth,
		authz:  authz,
		wshub:  wshub,
		router: router,
		store:  store,
		server: &http.Server{
			Addr:              fmt.Sprintf(":%v", cfg.Port),
			ReadHeaderTimeout: 5 * time.Second,
			Handler:           router.Handler(),
		},
	}

	api.registerRoutes()

	return api
}

func (api *API) Run(ctx context.Context) error {
	log.Info().Msgf("api - listening on port %v", api.config.Port)
	childCtx, cancel := context.WithCancel(ctx)
	var err error

	go func() {
		defer cancel()
		err = api.server.ListenAndServe()
	}()

	<-childCtx.Done()
	return err
}

func (api *API) Stop(ctx context.Context) error {
	return api.server.Shutdown(ctx)
}
