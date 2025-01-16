package budgets

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ndovnar/family-budget-api/internal/api/error"
	"github.com/ndovnar/family-budget-api/internal/model"
	"github.com/ndovnar/family-budget-api/internal/store"
	"github.com/rs/zerolog/log"
)

func (b *Budgets) HandleDeleteBudget(ctx *gin.Context) {
	id := ctx.Param("id")

	hasAccess := b.authz.IsUserHasAccessToBudget(ctx, id)
	if !hasAccess {
		ctx.Error(error.NewHttpError(http.StatusForbidden))
		return
	}

	budget, err := b.store.GetBudget(ctx, id)
	if err != nil {
		if err == store.ErrNotFound {
			ctx.Error(error.NewHttpError(http.StatusNotFound))
		} else {
			log.Error().Err(err).Msg("failed to get budget from store")
			ctx.Error(error.NewHttpError(http.StatusInternalServerError))
		}

		return
	}

	err = b.store.DeleteBudget(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg("failed to delete budget from store")
		ctx.Error(error.NewHttpError(http.StatusInternalServerError))
		return
	}

	b.broadcastUpdate(ctx, model.UpdateActionDelete, budget)
}
