package accounts

import (
	"context"

	"github.com/ndovnar/family-budget-api/internal/model"
)

func (a *Accounts) broadcastUpdate(ctx context.Context, updateAction model.UpdateActionType, account *model.Account) {
	userIDs := a.authz.GetUserIDsHaveAccessToAccount(ctx, account.ID)
	update := model.NewUpdate(model.UpdateTypeBudget, updateAction, account)

	a.wshub.Broadacst(update, userIDs...)
}
