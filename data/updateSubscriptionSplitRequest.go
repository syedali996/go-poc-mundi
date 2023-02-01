package data

type UpdateSubscriptionSplitRequest struct {
    Enabled bool                 `json:"enabled"`
    Rules   []CreateSplitRequest `json:"rules"`
}
