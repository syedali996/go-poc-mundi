package controllers

import (
    "fmt"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "pagarmeapisdk/data"
    "time"
)

type RecipientsController struct {
    baseController
}

func NewRecipientsController(baseController baseController) *RecipientsController {
    recipientsController := RecipientsController{baseController: baseController}
    return &recipientsController
}

// Updates a recipient
func (r *RecipientsController) UpdateRecipient(
    recipientId string,
    request data.UpdateRecipientRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetRecipientResponse],
    error) {
    req := r.prepareRequest("PUT", fmt.Sprintf("/recipients/%s", recipientId))
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetRecipientResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetRecipientResponse]{}, err
    }
    
    var result data.GetRecipientResponse
    result, err = utilities.DecodeResults[data.GetRecipientResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetRecipientResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Creates an anticipation
func (r *RecipientsController) CreateAnticipation(
    recipientId string,
    request data.CreateAnticipationRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetAnticipationResponse],
    error) {
    req := r.prepareRequest(
      "POST",
      fmt.Sprintf("/recipients/%s/anticipations", recipientId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetAnticipationResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetAnticipationResponse]{}, err
    }
    
    var result data.GetAnticipationResponse
    result, err = utilities.DecodeResults[data.GetAnticipationResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetAnticipationResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Gets the anticipation limits for a recipient
func (r *RecipientsController) GetAnticipationLimits(
    recipientId string,
    timeframe string,
    paymentDate time.Time) (
    https.ApiResponse[data.GetAnticipationLimitResponse],
    error) {
    req := r.prepareRequest(
      "GET",
      fmt.Sprintf("/recipients/%s/anticipation_limits", recipientId),
    )
    req.Authenticate(true)
    req.QueryParam("timeframe", timeframe)
    req.QueryParam("payment_date", paymentDate.Format(time.RFC3339))
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetAnticipationLimitResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetAnticipationLimitResponse]{}, err
    }
    
    var result data.GetAnticipationLimitResponse
    result, err = utilities.DecodeResults[data.GetAnticipationLimitResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetAnticipationLimitResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Retrieves paginated recipients information
func (r *RecipientsController) GetRecipients(
    page int,
    size int) (
    https.ApiResponse[data.ListRecipientResponse],
    error) {
    req := r.prepareRequest("GET", "/recipients")
    req.Authenticate(true)
    req.QueryParam("page", page)
    req.QueryParam("size", size)
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.ListRecipientResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.ListRecipientResponse]{}, err
    }
    
    var result data.ListRecipientResponse
    result, err = utilities.DecodeResults[data.ListRecipientResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.ListRecipientResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (r *RecipientsController) GetWithdrawById(
    recipientId string,
    withdrawalId string) (
    https.ApiResponse[data.GetWithdrawResponse],
    error) {
    req := r.prepareRequest(
      "GET",
      fmt.Sprintf("/recipients/%s/withdrawals/%s", recipientId, withdrawalId),
    )
    req.Authenticate(true)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetWithdrawResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetWithdrawResponse]{}, err
    }
    
    var result data.GetWithdrawResponse
    result, err = utilities.DecodeResults[data.GetWithdrawResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetWithdrawResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Updates the default bank account from a recipient
func (r *RecipientsController) UpdateRecipientDefaultBankAccount(
    recipientId string,
    request data.UpdateRecipientBankAccountRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetRecipientResponse],
    error) {
    req := r.prepareRequest(
      "PATCH",
      fmt.Sprintf("/recipients/%s/default-bank-account", recipientId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetRecipientResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetRecipientResponse]{}, err
    }
    
    var result data.GetRecipientResponse
    result, err = utilities.DecodeResults[data.GetRecipientResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetRecipientResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Updates recipient metadata
func (r *RecipientsController) UpdateRecipientMetadata(
    recipientId string,
    request data.UpdateMetadataRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetRecipientResponse],
    error) {
    req := r.prepareRequest(
      "PATCH",
      fmt.Sprintf("/recipients/%s/metadata", recipientId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetRecipientResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetRecipientResponse]{}, err
    }
    
    var result data.GetRecipientResponse
    result, err = utilities.DecodeResults[data.GetRecipientResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetRecipientResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Gets a paginated list of transfers for the recipient
func (r *RecipientsController) GetTransfers(
    recipientId string,
    page int,
    size int,
    status string,
    createdSince time.Time,
    createdUntil time.Time) (
    https.ApiResponse[data.ListTransferResponse],
    error) {
    req := r.prepareRequest(
      "GET",
      fmt.Sprintf("/recipients/%s/transfers", recipientId),
    )
    req.Authenticate(true)
    req.QueryParam("page", page)
    req.QueryParam("size", size)
    req.QueryParam("status", status)
    req.QueryParam("created_since", createdSince.Format(time.RFC3339))
    req.QueryParam("created_until", createdUntil.Format(time.RFC3339))
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.ListTransferResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.ListTransferResponse]{}, err
    }
    
    var result data.ListTransferResponse
    result, err = utilities.DecodeResults[data.ListTransferResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.ListTransferResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Gets a transfer
func (r *RecipientsController) GetTransfer(
    recipientId string,
    transferId string) (
    https.ApiResponse[data.GetTransferResponse],
    error) {
    req := r.prepareRequest(
      "GET",
      fmt.Sprintf("/recipients/%s/transfers/%s", recipientId, transferId),
    )
    req.Authenticate(true)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetTransferResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetTransferResponse]{}, err
    }
    
    var result data.GetTransferResponse
    result, err = utilities.DecodeResults[data.GetTransferResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetTransferResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (r *RecipientsController) CreateWithdraw(
    recipientId string,
    request data.CreateWithdrawRequest) (
    https.ApiResponse[data.GetWithdrawResponse],
    error) {
    req := r.prepareRequest(
      "POST",
      fmt.Sprintf("/recipients/%s/withdrawals", recipientId),
    )
    req.Authenticate(true)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetWithdrawResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetWithdrawResponse]{}, err
    }
    
    var result data.GetWithdrawResponse
    result, err = utilities.DecodeResults[data.GetWithdrawResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetWithdrawResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Updates recipient metadata
func (r *RecipientsController) UpdateAutomaticAnticipationSettings(
    recipientId string,
    request data.UpdateAutomaticAnticipationSettingsRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetRecipientResponse],
    error) {
    req := r.prepareRequest(
      "PATCH",
      fmt.Sprintf("/recipients/%s/automatic-anticipation-settings", recipientId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetRecipientResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetRecipientResponse]{}, err
    }
    
    var result data.GetRecipientResponse
    result, err = utilities.DecodeResults[data.GetRecipientResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetRecipientResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Gets an anticipation
func (r *RecipientsController) GetAnticipation(
    recipientId string,
    anticipationId string) (
    https.ApiResponse[data.GetAnticipationResponse],
    error) {
    req := r.prepareRequest(
      "GET",
      fmt.Sprintf("/recipients/%s/anticipations/%s", recipientId, anticipationId),
    )
    req.Authenticate(true)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetAnticipationResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetAnticipationResponse]{}, err
    }
    
    var result data.GetAnticipationResponse
    result, err = utilities.DecodeResults[data.GetAnticipationResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetAnticipationResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (r *RecipientsController) UpdateRecipientTransferSettings(
    recipientId string,
    request data.UpdateTransferSettingsRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetRecipientResponse],
    error) {
    req := r.prepareRequest(
      "PATCH",
      fmt.Sprintf("/recipients/%s/transfer-settings", recipientId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetRecipientResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetRecipientResponse]{}, err
    }
    
    var result data.GetRecipientResponse
    result, err = utilities.DecodeResults[data.GetRecipientResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetRecipientResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Retrieves a paginated list of anticipations from a recipient
func (r *RecipientsController) GetAnticipations(
    recipientId string,
    page int,
    size int,
    status string,
    timeframe string,
    paymentDateSince time.Time,
    paymentDateUntil time.Time,
    createdSince time.Time,
    createdUntil time.Time) (
    https.ApiResponse[data.ListAnticipationResponse],
    error) {
    req := r.prepareRequest(
      "GET",
      fmt.Sprintf("/recipients/%s/anticipations", recipientId),
    )
    req.Authenticate(true)
    req.QueryParam("page", page)
    req.QueryParam("size", size)
    req.QueryParam("status", status)
    req.QueryParam("timeframe", timeframe)
    req.QueryParam("payment_date_since", paymentDateSince.Format(time.RFC3339))
    req.QueryParam("payment_date_until", paymentDateUntil.Format(time.RFC3339))
    req.QueryParam("created_since", createdSince.Format(time.RFC3339))
    req.QueryParam("created_until", createdUntil.Format(time.RFC3339))
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.ListAnticipationResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.ListAnticipationResponse]{}, err
    }
    
    var result data.ListAnticipationResponse
    result, err = utilities.DecodeResults[data.ListAnticipationResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.ListAnticipationResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Retrieves recipient information
func (r *RecipientsController) GetRecipient(recipientId string) (
    https.ApiResponse[data.GetRecipientResponse],
    error) {
    req := r.prepareRequest("GET", fmt.Sprintf("/recipients/%s", recipientId))
    req.Authenticate(true)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetRecipientResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetRecipientResponse]{}, err
    }
    
    var result data.GetRecipientResponse
    result, err = utilities.DecodeResults[data.GetRecipientResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetRecipientResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Get balance information for a recipient
func (r *RecipientsController) GetBalance(recipientId string) (
    https.ApiResponse[data.GetBalanceResponse],
    error) {
    req := r.prepareRequest(
      "GET",
      fmt.Sprintf("/recipients/%s/balance", recipientId),
    )
    req.Authenticate(true)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetBalanceResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetBalanceResponse]{}, err
    }
    
    var result data.GetBalanceResponse
    result, err = utilities.DecodeResults[data.GetBalanceResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetBalanceResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Gets a paginated list of transfers for the recipient
func (r *RecipientsController) GetWithdrawals(
    recipientId string,
    page int,
    size int,
    status string,
    createdSince time.Time,
    createdUntil time.Time) (
    https.ApiResponse[data.ListWithdrawals],
    error) {
    req := r.prepareRequest(
      "GET",
      fmt.Sprintf("/recipients/%s/withdrawals", recipientId),
    )
    req.Authenticate(true)
    req.QueryParam("page", page)
    req.QueryParam("size", size)
    req.QueryParam("status", status)
    req.QueryParam("created_since", createdSince.Format(time.RFC3339))
    req.QueryParam("created_until", createdUntil.Format(time.RFC3339))
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.ListWithdrawals]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.ListWithdrawals]{}, err
    }
    
    var result data.ListWithdrawals
    result, err = utilities.DecodeResults[data.ListWithdrawals](decoder)
    if err != nil {
        return https.ApiResponse[data.ListWithdrawals]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Creates a transfer for a recipient
func (r *RecipientsController) CreateTransfer(
    recipientId string,
    request data.CreateTransferRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetTransferResponse],
    error) {
    req := r.prepareRequest(
      "POST",
      fmt.Sprintf("/recipients/%s/transfers", recipientId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetTransferResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetTransferResponse]{}, err
    }
    
    var result data.GetTransferResponse
    result, err = utilities.DecodeResults[data.GetTransferResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetTransferResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Creates a new recipient
func (r *RecipientsController) CreateRecipient(
    request data.CreateRecipientRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetRecipientResponse],
    error) {
    req := r.prepareRequest("POST", "/recipients")
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetRecipientResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetRecipientResponse]{}, err
    }
    
    var result data.GetRecipientResponse
    result, err = utilities.DecodeResults[data.GetRecipientResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetRecipientResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Retrieves recipient information
func (r *RecipientsController) GetRecipientByCode(code string) (
    https.ApiResponse[data.GetRecipientResponse],
    error) {
    req := r.prepareRequest("GET", fmt.Sprintf("/recipients/%s", code))
    req.Authenticate(true)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetRecipientResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetRecipientResponse]{}, err
    }
    
    var result data.GetRecipientResponse
    result, err = utilities.DecodeResults[data.GetRecipientResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetRecipientResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (r *RecipientsController) GetDefaultRecipient() (
    https.ApiResponse[data.GetRecipientResponse],
    error) {
    req := r.prepareRequest("GET", "/recipients/default")
    req.Authenticate(true)
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetRecipientResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetRecipientResponse]{}, err
    }
    
    var result data.GetRecipientResponse
    result, err = utilities.DecodeResults[data.GetRecipientResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetRecipientResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}
