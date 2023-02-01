package data

// Request for creating a price bracket
type CreatePriceBracketRequest struct {
    StartQuantity int `json:"start_quantity"`
    Price         int `json:"price"`
    EndQuantity   int `json:"end_quantity,omitempty"`
    OveragePrice  int `json:"overage_price,omitempty"`
}
