package data

// Response model for listing subscription items
type ListSubscriptionItemsResponse struct {
    Data   *[]GetSubscriptionItemResponse `json:"data"`
    Paging *PagingResponse                `json:"paging"`
}
