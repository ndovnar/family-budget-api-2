package budgets

import "github.com/ndovnar/family-budget-api/internal/model"

type budgetsResponse struct {
	Values []*model.Budget `json:"values"`
	Meta   *meta
}

type meta struct {
	Count int64 `json:"count"`
}

func newBudgetsResponse(budgets []*model.Budget, count int64) *budgetsResponse {
	return &budgetsResponse{
		Values: budgets,
		Meta: &meta{
			Count: count,
		},
	}
}
