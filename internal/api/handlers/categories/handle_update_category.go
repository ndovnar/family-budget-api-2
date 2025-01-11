package categories

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ndovnar/family-budget-api/internal/api/error"
	"github.com/ndovnar/family-budget-api/internal/model"
	"github.com/rs/zerolog/log"
)

func (c *Categories) HandleUpdateCategory(ctx *gin.Context) {
	id := ctx.Param("id")

	hasAccess := c.authz.IsUserHasAccessToCategory(ctx, id)
	if !hasAccess {
		ctx.Error(error.NewHttpError(http.StatusForbidden))
		return
	}

	var req categoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(error.NewHttpError(http.StatusBadRequest))
		return
	}

	category, err := c.store.UpdateCategory(ctx, id, &model.Category{
		BudgetID: req.BudgetID,
		Name:     req.Name,
		Currency: req.Currency,
		Balance:  req.Balance,
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to updated category")
		ctx.Error(error.NewHttpError(http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, category)
}
