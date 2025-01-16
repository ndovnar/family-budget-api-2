package transactions

import (
	"context"

	"github.com/AlekSi/pointer"
	"github.com/ndovnar/family-budget-api/internal/model"
)

func (t *Transactions) broadcastUpdate(ctx context.Context, updateAction model.UpdateActionType, transaction *model.Transaction) {
	userIDs := []string{}

	if transaction.FromAccountID != nil {
		fromAccountUserIDs := t.authz.GetUserIDsHaveAccessToAccount(ctx, pointer.GetString(transaction.FromAccountID))
		userIDs = append(userIDs, fromAccountUserIDs...)
	}

	if transaction.ToAccountID != nil {
		toAccountUserIDs := t.authz.GetUserIDsHaveAccessToAccount(ctx, pointer.GetString(transaction.FromAccountID))
		userIDs = append(userIDs, toAccountUserIDs...)
	}

	update := model.NewUpdate(model.UpdateTypeTransaction, updateAction, transaction)
	t.wshub.Broadacst(update, userIDs...)
}
