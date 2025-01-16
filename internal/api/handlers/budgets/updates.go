package budgets

import (
	"context"

	"github.com/ndovnar/family-budget-api/internal/model"
)

func (b *Budgets) broadcastUpdate(ctx context.Context, updateAction model.UpdateActionType, budget *model.Budget) {
	userIDs := b.authz.GetUserIDsHaveAccessToBudget(ctx, budget.ID)
	update := model.NewUpdate(model.UpdateTypeBudget, updateAction, budget)

	b.wshub.Broadacst(update, userIDs...)
}
