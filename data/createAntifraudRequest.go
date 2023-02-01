package data

type CreateAntifraudRequest struct {
    Type      string                 `json:"type"`
    Clearsale CreateClearSaleRequest `json:"clearsale"`
}
