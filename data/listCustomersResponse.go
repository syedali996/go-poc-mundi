package data

// Response for listing the customers
type ListCustomersResponse struct {
    Data   *[]GetCustomerResponse `json:"data"`
    Paging *PagingResponse        `json:"paging"`
}
