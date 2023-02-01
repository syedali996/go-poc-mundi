package data

// Anticipation limits
type GetAnticipationLimitsResponse struct {
    Max *GetAnticipationLimitResponse `json:"max"`
    Min *GetAnticipationLimitResponse `json:"min"`
}
