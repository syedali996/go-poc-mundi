package data

// Token data
type CreateTokenRequest struct {
    Type string                 `json:"type"`
    Card CreateCardTokenRequest `json:"card"`
}
