package authz

import (
	"context"

	"github.com/gin-gonic/gin"
)

func (a *Authz) IsUserHasAccessToCategory(ctx *gin.Context, id string) bool {
	category, err := a.store.GetCategory(ctx, id)
	if err != nil {
		return false
	}

	return a.IsUserHasAccessToBudget(ctx, category.BudgetID)
}

func (a *Authz) GetUserIDsHaveAccessToCategory(ctx context.Context, id string) []string {
	category, err := a.store.GetCategory(ctx, id)
	if err != nil {
		return []string{}
	}

	return a.GetUserIDsHaveAccessToBudget(ctx, category.BudgetID)
}
