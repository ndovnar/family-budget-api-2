package accounts

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ndovnar/family-budget-api/internal/api/error"
	"github.com/ndovnar/family-budget-api/internal/model"
	"github.com/rs/zerolog/log"
)

func (a *Accounts) HandleCreateAccount(ctx *gin.Context) {
	claims := a.auth.GetClaimsFromContext(ctx)

	var req accountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(error.NewHttpError(http.StatusBadRequest))
		return
	}

	account, err := a.store.CreateAccount(ctx, &model.Account{
		OwnerID: claims.UserID,
		Name:    req.Name,
		Balance: req.Balance,
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to create account")
		ctx.Error(error.NewHttpError(http.StatusInternalServerError))
		return
	}

	a.broadcastUpdate(ctx, model.UpdateActionCreate, account)
	ctx.JSON(http.StatusOK, account)
}
