package model

type UpdateActionType string
type UpdateType string

const (
	UpdateActionCreate UpdateActionType = "create"
	UpdateActionUpdate UpdateActionType = "update"
	UpdateActionDelete UpdateActionType = "delete"

	UpdateTypeAccount     UpdateType = "account"
	UpdateTypeBudget      UpdateType = "budget"
	UpdateTypeCategory    UpdateType = "category"
	UpdateTypeTransaction UpdateType = "transaction"
)

type Update[T any] struct {
	Type   UpdateType       `json:"type"`
	Action UpdateActionType `json:"action"`
	Data   T                `json:"data"`
}

func NewUpdate[T any](updateType UpdateType, action UpdateActionType, data T) *Update[T] {
	return &Update[T]{
		Type:   updateType,
		Action: action,
		Data:   data,
	}
}
