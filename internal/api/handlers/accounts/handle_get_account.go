package accounts

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ndovnar/family-budget-api/internal/api/error"
	"github.com/ndovnar/family-budget-api/internal/store"
)

func (a *Accounts) HandleGetAcccount(ctx *gin.Context) {
	id := ctx.Param("id")

	hasAccess := a.authz.IsUserHasAccessToAccount(ctx, id)
	if !hasAccess {
		ctx.Error(error.NewHttpError(http.StatusForbidden))
		return
	}

	account, err := a.store.GetAccount(ctx, id)

	if err != nil {
		if err == store.ErrNotFound {
			ctx.Error(error.NewHttpError(http.StatusNotFound))
		} else {
			ctx.Error(error.NewHttpError(http.StatusInternalServerError))
		}

		return
	}

	ctx.JSON(http.StatusOK, account)
}
