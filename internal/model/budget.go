package model

type Budget struct {
	ID        string `json:"id,omitempty" bson:"_id,omitempty"`
	Owner     string `json:"owner" bson:"owner"`
	Name      string `json:"name" bson:"name"`
	IsDeleted bool   `json:"deleted,omitempty" bson:"deleted"`
	Dates     Dates  `json:"dates" bson:"dates"`
}

type GetBudgetsFilter struct {
	Owner   string
	Deleted bool
}
