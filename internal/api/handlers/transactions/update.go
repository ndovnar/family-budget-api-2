package transactions

import (
	"context"

	"github.com/ndovnar/family-budget-api/internal/model"
)

func (t *Transactions) broadcastUpdate(ctx context.Context, updateAction model.UpdateActionType, transaction *model.Transaction) {
	userIDs := []string{}

	if transaction.FromAccountID != "" {
		fromAccountUserIDs := t.authz.GetUserIDsHaveAccessToAccount(ctx, transaction.FromAccountID)
		userIDs = append(userIDs, fromAccountUserIDs...)
	}

	if transaction.ToAccountID != "" {
		toAccountUserIDs := t.authz.GetUserIDsHaveAccessToAccount(ctx, transaction.FromAccountID)
		userIDs = append(userIDs, toAccountUserIDs...)
	}

	update := model.NewUpdate(model.UpdateTypeTransaction, updateAction, transaction)
	t.wshub.Broadacst(update, userIDs...)
}
