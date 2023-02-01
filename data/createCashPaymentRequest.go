package data

type CreateCashPaymentRequest struct {
    Description string `json:"description"`
    Confirm     bool   `json:"confirm"`
}
