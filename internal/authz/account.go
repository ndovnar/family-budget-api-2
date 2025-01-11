package authz

import (
	"context"

	"github.com/gin-gonic/gin"
)

func (a *Authz) IsUserHasAccessToAccount(ctx *gin.Context, id string) bool {
	claims := a.auth.GetClaimsFromContext(ctx)

	account, err := a.store.GetAccount(ctx, id)
	if err != nil {
		return false
	}

	return account.OwnerID == claims.UserID
}

func (a *Authz) GetUserIDsHaveAccessToAccount(ctx context.Context, id string) []string {
	account, err := a.store.GetAccount(ctx, id)
	if err != nil {
		return []string{}
	}

	return []string{account.OwnerID}
}
