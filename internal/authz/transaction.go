package authz

import (
	"github.com/AlekSi/pointer"
	"github.com/gin-gonic/gin"
)

func (a *Authz) IsUserHasReadAcessToTransaction(ctx *gin.Context, id string) bool {
	transaction, err := a.store.GetTransaction(ctx, id)
	if err != nil {
		return false
	}

	if transaction.CategoryID == nil {
		return false
	}

	return a.IsUserHasAccessToCategory(ctx, pointer.GetString(transaction.CategoryID))
}

func (a *Authz) IsUserHasWriteAcessToTransaction(ctx *gin.Context, id string) bool {
	claims := a.auth.GetClaimsFromContext(ctx)

	transaction, err := a.store.GetTransaction(ctx, id)
	if err != nil {
		return false
	}

	return transaction.UserID == claims.UserID
}
