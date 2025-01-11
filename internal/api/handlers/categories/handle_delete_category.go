package categories

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ndovnar/family-budget-api/internal/api/error"
	"github.com/ndovnar/family-budget-api/internal/model"
	"github.com/ndovnar/family-budget-api/internal/store"
	"github.com/rs/zerolog/log"
)

func (c *Categories) HandleDeleteCategory(ctx *gin.Context) {
	id := ctx.Param("id")

	hasAccess := c.authz.IsUserHasAccessToCategory(ctx, id)
	if !hasAccess {
		ctx.Error(error.NewHttpError(http.StatusForbidden))
		return
	}

	category, err := c.store.GetCategory(ctx, id)
	if err != nil {
		if err == store.ErrNotFound {
			ctx.Error(error.NewHttpError(http.StatusNotFound))
		} else {
			log.Error().Err(err).Msg("failed to get category from store")
			ctx.Error(error.NewHttpError(http.StatusInternalServerError))
		}

		return
	}

	err = c.store.DeleteCategory(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg("failed to delete category from store")
		ctx.Error(error.NewHttpError(http.StatusInternalServerError))
		return
	}

	c.broadcastUpdate(ctx, model.UpdateActionDelete, category)
}
