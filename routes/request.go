package routes

type TransferRequest struct {
	AccountFrom int64 `json:"account_from" validate:"required"`
	AccountTo   int64 `json:"account_to" validate:"required"`
	Amount      int64 `json:"amount" validate:"required"`
}
