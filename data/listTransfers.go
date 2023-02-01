package data

type ListTransfers struct {
    Data   []GetTransfer  `json:"data"`
    Paging PagingResponse `json:"paging"`
}
