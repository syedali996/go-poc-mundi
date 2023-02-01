package controllers

import (
    "fmt"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "pagarmeapisdk/data"
    "time"
)

type ChargesController struct {
    baseController
}

func NewChargesController(baseController baseController) *ChargesController {
    chargesController := ChargesController{baseController: baseController}
    return &chargesController
}

// Updates the metadata from a charge
func (c *ChargesController) UpdateChargeMetadata(
    chargeId string,
    request data.UpdateMetadataRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetChargeResponse],
    error) {
    req := c.prepareRequest("PATCH", fmt.Sprintf("/Charges/%s/metadata", chargeId))
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetChargeResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetChargeResponse]{}, err
    }
    
    var result data.GetChargeResponse
    result, err = utilities.DecodeResults[data.GetChargeResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetChargeResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Updates a charge's payment method
func (c *ChargesController) UpdateChargePaymentMethod(
    chargeId string,
    request data.UpdateChargePaymentMethodRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetChargeResponse],
    error) {
    req := c.prepareRequest(
      "PATCH",
      fmt.Sprintf("/charges/%s/payment-method", chargeId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetChargeResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetChargeResponse]{}, err
    }
    
    var result data.GetChargeResponse
    result, err = utilities.DecodeResults[data.GetChargeResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetChargeResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (c *ChargesController) GetChargeTransactions(
    chargeId string,
    page int,
    size int) (
    https.ApiResponse[data.ListChargeTransactionsResponse],
    error) {
    req := c.prepareRequest(
      "GET",
      fmt.Sprintf("/charges/%s/transactions", chargeId),
    )
    req.Authenticate(true)
    req.QueryParam("page", page)
    req.QueryParam("size", size)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.ListChargeTransactionsResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.ListChargeTransactionsResponse]{}, err
    }
    
    var result data.ListChargeTransactionsResponse
    result, err = utilities.DecodeResults[data.ListChargeTransactionsResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.ListChargeTransactionsResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Updates the due date from a charge
func (c *ChargesController) UpdateChargeDueDate(
    chargeId string,
    request data.UpdateChargeDueDateRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetChargeResponse],
    error) {
    req := c.prepareRequest("PATCH", fmt.Sprintf("/Charges/%s/due-date", chargeId))
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetChargeResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetChargeResponse]{}, err
    }
    
    var result data.GetChargeResponse
    result, err = utilities.DecodeResults[data.GetChargeResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetChargeResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Lists all charges
func (c *ChargesController) GetCharges(
    page int,
    size int,
    code string,
    status string,
    paymentMethod string,
    customerId string,
    orderId string,
    createdSince time.Time,
    createdUntil time.Time) (
    https.ApiResponse[data.ListChargesResponse],
    error) {
    req := c.prepareRequest("GET", "/charges")
    req.Authenticate(true)
    req.QueryParam("page", page)
    req.QueryParam("size", size)
    req.QueryParam("code", code)
    req.QueryParam("status", status)
    req.QueryParam("payment_method", paymentMethod)
    req.QueryParam("customer_id", customerId)
    req.QueryParam("order_id", orderId)
    req.QueryParam("created_since", createdSince.Format(time.RFC3339))
    req.QueryParam("created_until", createdUntil.Format(time.RFC3339))
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.ListChargesResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.ListChargesResponse]{}, err
    }
    
    var result data.ListChargesResponse
    result, err = utilities.DecodeResults[data.ListChargesResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.ListChargesResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Captures a charge
func (c *ChargesController) CaptureCharge(
    chargeId string,
    request data.CreateCaptureChargeRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetChargeResponse],
    error) {
    req := c.prepareRequest("POST", fmt.Sprintf("/charges/%s/capture", chargeId))
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetChargeResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetChargeResponse]{}, err
    }
    
    var result data.GetChargeResponse
    result, err = utilities.DecodeResults[data.GetChargeResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetChargeResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Updates the card from a charge
func (c *ChargesController) UpdateChargeCard(
    chargeId string,
    request data.UpdateChargeCardRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetChargeResponse],
    error) {
    req := c.prepareRequest("PATCH", fmt.Sprintf("/charges/%s/card", chargeId))
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetChargeResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetChargeResponse]{}, err
    }
    
    var result data.GetChargeResponse
    result, err = utilities.DecodeResults[data.GetChargeResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetChargeResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Get a charge from its id
func (c *ChargesController) GetCharge(chargeId string) (
    https.ApiResponse[data.GetChargeResponse],
    error) {
    req := c.prepareRequest("GET", fmt.Sprintf("/charges/%s", chargeId))
    req.Authenticate(true)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetChargeResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetChargeResponse]{}, err
    }
    
    var result data.GetChargeResponse
    result, err = utilities.DecodeResults[data.GetChargeResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetChargeResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (c *ChargesController) GetChargesSummary(
    status string,
    createdSince time.Time,
    createdUntil time.Time) (
    https.ApiResponse[data.GetChargesSummaryResponse],
    error) {
    req := c.prepareRequest("GET", "/charges/summary")
    req.Authenticate(true)
    req.QueryParam("status", status)
    req.QueryParam("created_since", createdSince.Format(time.RFC3339))
    req.QueryParam("created_until", createdUntil.Format(time.RFC3339))
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetChargesSummaryResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetChargesSummaryResponse]{}, err
    }
    
    var result data.GetChargesSummaryResponse
    result, err = utilities.DecodeResults[data.GetChargesSummaryResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetChargesSummaryResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Retries a charge
func (c *ChargesController) RetryCharge(
    chargeId string,
    idempotencyKey string) (
    https.ApiResponse[data.GetChargeResponse],
    error) {
    req := c.prepareRequest("POST", fmt.Sprintf("/charges/%s/retry", chargeId))
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetChargeResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetChargeResponse]{}, err
    }
    
    var result data.GetChargeResponse
    result, err = utilities.DecodeResults[data.GetChargeResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetChargeResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Cancel a charge
func (c *ChargesController) CancelCharge(
    chargeId string,
    request data.CreateCancelChargeRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetChargeResponse],
    error) {
    req := c.prepareRequest("DELETE", fmt.Sprintf("/charges/%s", chargeId))
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetChargeResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetChargeResponse]{}, err
    }
    
    var result data.GetChargeResponse
    result, err = utilities.DecodeResults[data.GetChargeResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetChargeResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Creates a new charge
func (c *ChargesController) CreateCharge(
    request data.CreateChargeRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetChargeResponse],
    error) {
    req := c.prepareRequest("POST", "/Charges")
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetChargeResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetChargeResponse]{}, err
    }
    
    var result data.GetChargeResponse
    result, err = utilities.DecodeResults[data.GetChargeResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetChargeResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (c *ChargesController) ConfirmPayment(
    chargeId string,
    request data.CreateConfirmPaymentRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetChargeResponse],
    error) {
    req := c.prepareRequest(
      "POST",
      fmt.Sprintf("/charges/%s/confirm-payment", chargeId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetChargeResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetChargeResponse]{}, err
    }
    
    var result data.GetChargeResponse
    result, err = utilities.DecodeResults[data.GetChargeResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetChargeResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}
