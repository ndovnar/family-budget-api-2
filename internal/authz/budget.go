package authz

import (
	"context"

	"github.com/gin-gonic/gin"
)

func (a *Authz) IsUserHasAccessToBudget(ctx *gin.Context, id string) bool {
	claims := a.auth.GetClaimsFromContext(ctx)

	budget, err := a.store.GetBudget(ctx, id)
	if err != nil {
		return false
	}

	return budget.OwnerID == claims.UserID
}

func (a *Authz) GetUserIDsHaveAccessToBudget(ctx context.Context, id string) []string {
	budget, err := a.store.GetBudget(ctx, id)
	if err != nil {
		return []string{}
	}

	return []string{budget.OwnerID}
}
