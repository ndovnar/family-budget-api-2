package transactions

const ColName = "name”"

type transactionRequest struct {
	FromAccountID string  `json:"fromAccount" binding:"required"`
	ToAccountID   string  `json:"toAccount" binding:"required"`
	CategoryID    string  `json:"category"`
	Amount        float64 `json:"amount"`
	Description   string  `json:"description"`
}
