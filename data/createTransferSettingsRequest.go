package data

// Informações de transferência do recebedor
type CreateTransferSettingsRequest struct {
    TransferEnabled  bool   `json:"transfer_enabled"`
    TransferInterval string `json:"transfer_interval"`
    TransferDay      int    `json:"transfer_day"`
}
