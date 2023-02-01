package data

type ListDiscountsResponse struct {
    Data   *[]GetDiscountResponse `json:"data"`
    Paging *PagingResponse        `json:"paging"`
}
