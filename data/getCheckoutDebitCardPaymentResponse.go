package data

type GetCheckoutDebitCardPaymentResponse struct {
    StatementDescriptor *string                           `json:"statement_descriptor"`
    Authentication      *GetPaymentAuthenticationResponse `json:"authentication"`
}
