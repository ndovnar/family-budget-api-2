package transactions

import (
	"net/http"

	"github.com/AlekSi/pointer"
	"github.com/gin-gonic/gin"
	"github.com/ndovnar/family-budget-api/internal/api/error"
	"github.com/ndovnar/family-budget-api/internal/model"
	"github.com/rs/zerolog/log"
)

func (t *Transactions) HandleUpdateTransaction(ctx *gin.Context) {
	id := ctx.Param("id")

	hasAccess := t.authz.IsUserHasWriteAcessToTransaction(ctx, id)
	if !hasAccess {
		ctx.Error(error.NewHttpError(http.StatusForbidden))
		return
	}

	var req transactionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(error.NewHttpError(http.StatusBadRequest))
		return
	}

	updatedTransaction, err := t.store.UpdateTransaction(ctx, id, &model.Transaction{
		FromAccountID: pointer.ToStringOrNil(req.FromAccountID),
		ToAccountID:   pointer.ToStringOrNil(req.ToAccountID),
		CategoryID:    pointer.ToStringOrNil(req.CategoryID),
		Amount:        req.Amount,
		Description:   req.Description,
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to updated transaction")
		ctx.Error(error.NewHttpError(http.StatusInternalServerError))
		return
	}

	t.broadcastUpdate(ctx, model.UpdateActionUpdate, updatedTransaction)
	ctx.JSON(http.StatusOK, updatedTransaction)
}
