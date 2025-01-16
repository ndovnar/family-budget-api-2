package tokens

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ndovnar/family-budget-api/internal/api/error"
	"github.com/rs/zerolog/log"
)

func (t *Tokens) HandleWsAccessToken(ctx *gin.Context) {
	claims := t.auth.GetClaimsFromContext(ctx)

	wsAccessToken, err := t.auth.CreateWsAccessToken(claims.UserID)
	if err != nil {
		log.Error().Err(err).Msg("failed to create ws access token")
		ctx.Error(error.NewHttpError(http.StatusInternalServerError))
		return
	}

	resp := newWsAccessTokenResponse(wsAccessToken)
	ctx.JSON(http.StatusOK, resp)
}
