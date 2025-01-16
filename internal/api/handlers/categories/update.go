package categories

import (
	"context"

	"github.com/ndovnar/family-budget-api/internal/model"
)

func (t *Categories) broadcastUpdate(ctx context.Context, updateAction model.UpdateActionType, category *model.Category) {
	userIDs := t.authz.GetUserIDsHaveAccessToCategory(ctx, category.ID)
	update := model.NewUpdate(model.UpdateTypeCategory, updateAction, category)
	t.wshub.Broadacst(update, userIDs...)
}
