package controllers

import (
    "fmt"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "net/http"
    "pagarmeapisdk/data"
    "time"
)

type SubscriptionsController struct {
    baseController
}

func NewSubscriptionsController(baseController baseController) *SubscriptionsController {
    subscriptionsController := SubscriptionsController{baseController: baseController}
    return &subscriptionsController
}

// TODO: type endpoint description here
func (s *SubscriptionsController) RenewSubscription(
    subscriptionId string,
    idempotencyKey string) (
    https.ApiResponse[data.GetPeriodResponse],
    error) {
    req := s.prepareRequest(
      "POST",
      fmt.Sprintf("/subscriptions/%s/cycles", subscriptionId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetPeriodResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetPeriodResponse]{}, err
    }
    
    var result data.GetPeriodResponse
    result, err = utilities.DecodeResults[data.GetPeriodResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetPeriodResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Updates the credit card from a subscription
func (s *SubscriptionsController) UpdateSubscriptionCard(
    subscriptionId string,
    request data.UpdateSubscriptionCardRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetSubscriptionResponse],
    error) {
    req := s.prepareRequest(
      "PATCH",
      fmt.Sprintf("/subscriptions/%s/card", subscriptionId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionResponse]{}, err
    }
    
    var result data.GetSubscriptionResponse
    result, err = utilities.DecodeResults[data.GetSubscriptionResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Deletes a usage
func (s *SubscriptionsController) DeleteUsage(
    subscriptionId string,
    itemId string,
    usageId string,
    idempotencyKey string) (
    https.ApiResponse[data.GetUsageResponse],
    error) {
    req := s.prepareRequest(
      "DELETE",
      fmt.Sprintf("/subscriptions/%s/items/%s/usages/%s", subscriptionId, itemId, usageId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetUsageResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetUsageResponse]{}, err
    }
    
    var result data.GetUsageResponse
    result, err = utilities.DecodeResults[data.GetUsageResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetUsageResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Creates a discount
func (s *SubscriptionsController) CreateDiscount(
    subscriptionId string,
    request data.CreateDiscountRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetDiscountResponse],
    error) {
    req := s.prepareRequest(
      "POST",
      fmt.Sprintf("/subscriptions/%s/discounts", subscriptionId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetDiscountResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetDiscountResponse]{}, err
    }
    
    var result data.GetDiscountResponse
    result, err = utilities.DecodeResults[data.GetDiscountResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetDiscountResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Create Usage
func (s *SubscriptionsController) CreateAnUsage(
    subscriptionId string,
    itemId string,
    idempotencyKey string) (
    https.ApiResponse[data.GetUsageResponse],
    error) {
    req := s.prepareRequest(
      "POST",
      fmt.Sprintf("/subscriptions/%s/items/%s/usages", subscriptionId, itemId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetUsageResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetUsageResponse]{}, err
    }
    
    var result data.GetUsageResponse
    result, err = utilities.DecodeResults[data.GetUsageResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetUsageResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (s *SubscriptionsController) UpdateCurrentCycleStatus(
    subscriptionId string,
    request data.UpdateCurrentCycleStatusRequest,
    idempotencyKey string) (
    *http.Response,
    error) {
    req := s.prepareRequest(
      "PATCH",
      fmt.Sprintf("/subscriptions/%s/cycle-status", subscriptionId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    context, err := req.Call()
    err = validateResponse(*context.Response)
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}

// Deletes a discount
func (s *SubscriptionsController) DeleteDiscount(
    subscriptionId string,
    discountId string,
    idempotencyKey string) (
    https.ApiResponse[data.GetDiscountResponse],
    error) {
    req := s.prepareRequest(
      "DELETE",
      fmt.Sprintf("/subscriptions/%s/discounts/%s", subscriptionId, discountId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetDiscountResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetDiscountResponse]{}, err
    }
    
    var result data.GetDiscountResponse
    result, err = utilities.DecodeResults[data.GetDiscountResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetDiscountResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Get Subscription Items
func (s *SubscriptionsController) GetSubscriptionItems(
    subscriptionId string,
    page int,
    size int,
    name string,
    code string,
    status string,
    description string,
    createdSince string,
    createdUntil string) (
    https.ApiResponse[data.ListSubscriptionItemsResponse],
    error) {
    req := s.prepareRequest(
      "GET",
      fmt.Sprintf("/subscriptions/%s/items", subscriptionId),
    )
    req.Authenticate(true)
    req.QueryParam("page", page)
    req.QueryParam("size", size)
    req.QueryParam("name", name)
    req.QueryParam("code", code)
    req.QueryParam("status", status)
    req.QueryParam("description", description)
    req.QueryParam("created_since", createdSince)
    req.QueryParam("created_until", createdUntil)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.ListSubscriptionItemsResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.ListSubscriptionItemsResponse]{}, err
    }
    
    var result data.ListSubscriptionItemsResponse
    result, err = utilities.DecodeResults[data.ListSubscriptionItemsResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.ListSubscriptionItemsResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Updates the payment method from a subscription
func (s *SubscriptionsController) UpdateSubscriptionPaymentMethod(
    subscriptionId string,
    request data.UpdateSubscriptionPaymentMethodRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetSubscriptionResponse],
    error) {
    req := s.prepareRequest(
      "PATCH",
      fmt.Sprintf("/subscriptions/%s/payment-method", subscriptionId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionResponse]{}, err
    }
    
    var result data.GetSubscriptionResponse
    result, err = utilities.DecodeResults[data.GetSubscriptionResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Get Subscription Item
func (s *SubscriptionsController) GetSubscriptionItem(
    subscriptionId string,
    itemId string) (
    https.ApiResponse[data.GetSubscriptionItemResponse],
    error) {
    req := s.prepareRequest(
      "GET",
      fmt.Sprintf("/subscriptions/%s/items/%s", subscriptionId, itemId),
    )
    req.Authenticate(true)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionItemResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionItemResponse]{}, err
    }
    
    var result data.GetSubscriptionItemResponse
    result, err = utilities.DecodeResults[data.GetSubscriptionItemResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionItemResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Gets all subscriptions
func (s *SubscriptionsController) GetSubscriptions(
    page int,
    size int,
    code string,
    billingType string,
    customerId string,
    planId string,
    cardId string,
    status string,
    nextBillingSince time.Time,
    nextBillingUntil time.Time,
    createdSince time.Time,
    createdUntil time.Time) (
    https.ApiResponse[data.ListSubscriptionsResponse],
    error) {
    req := s.prepareRequest("GET", "/subscriptions")
    req.Authenticate(true)
    req.QueryParam("page", page)
    req.QueryParam("size", size)
    req.QueryParam("code", code)
    req.QueryParam("billing_type", billingType)
    req.QueryParam("customer_id", customerId)
    req.QueryParam("plan_id", planId)
    req.QueryParam("card_id", cardId)
    req.QueryParam("status", status)
    req.QueryParam("next_billing_since", nextBillingSince.Format(time.RFC3339))
    req.QueryParam("next_billing_until", nextBillingUntil.Format(time.RFC3339))
    req.QueryParam("created_since", createdSince.Format(time.RFC3339))
    req.QueryParam("created_until", createdUntil.Format(time.RFC3339))
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.ListSubscriptionsResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.ListSubscriptionsResponse]{}, err
    }
    
    var result data.ListSubscriptionsResponse
    result, err = utilities.DecodeResults[data.ListSubscriptionsResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.ListSubscriptionsResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Cancels a subscription
func (s *SubscriptionsController) CancelSubscription(
    subscriptionId string,
    request data.CreateCancelSubscriptionRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetSubscriptionResponse],
    error) {
    req := s.prepareRequest(
      "DELETE",
      fmt.Sprintf("/subscriptions/%s", subscriptionId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionResponse]{}, err
    }
    
    var result data.GetSubscriptionResponse
    result, err = utilities.DecodeResults[data.GetSubscriptionResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Creates a increment
func (s *SubscriptionsController) CreateIncrement(
    subscriptionId string,
    request data.CreateIncrementRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetIncrementResponse],
    error) {
    req := s.prepareRequest(
      "POST",
      fmt.Sprintf("/subscriptions/%s/increments", subscriptionId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetIncrementResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetIncrementResponse]{}, err
    }
    
    var result data.GetIncrementResponse
    result, err = utilities.DecodeResults[data.GetIncrementResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetIncrementResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Creates a usage
func (s *SubscriptionsController) CreateUsage(
    subscriptionId string,
    itemId string,
    body data.CreateUsageRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetUsageResponse],
    error) {
    req := s.prepareRequest(
      "POST",
      fmt.Sprintf("/subscriptions/%s/items/%s/usages", subscriptionId, itemId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(body)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetUsageResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetUsageResponse]{}, err
    }
    
    var result data.GetUsageResponse
    result, err = utilities.DecodeResults[data.GetUsageResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetUsageResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (s *SubscriptionsController) GetDiscountById(
    subscriptionId string,
    discountId string) (
    https.ApiResponse[data.GetDiscountResponse],
    error) {
    req := s.prepareRequest(
      "GET",
      fmt.Sprintf("/subscriptions/%s/discounts/%s", subscriptionId, discountId),
    )
    req.Authenticate(true)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetDiscountResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetDiscountResponse]{}, err
    }
    
    var result data.GetDiscountResponse
    result, err = utilities.DecodeResults[data.GetDiscountResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetDiscountResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Creates a new subscription
func (s *SubscriptionsController) CreateSubscription(
    body data.CreateSubscriptionRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetSubscriptionResponse],
    error) {
    req := s.prepareRequest("POST", "/subscriptions")
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(body)
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionResponse]{}, err
    }
    
    var result data.GetSubscriptionResponse
    result, err = utilities.DecodeResults[data.GetSubscriptionResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (s *SubscriptionsController) GetIncrementById(
    subscriptionId string,
    incrementId string) (
    https.ApiResponse[data.GetIncrementResponse],
    error) {
    req := s.prepareRequest(
      "GET",
      fmt.Sprintf("/subscriptions/%s/increments/%s", subscriptionId, incrementId),
    )
    req.Authenticate(true)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetIncrementResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetIncrementResponse]{}, err
    }
    
    var result data.GetIncrementResponse
    result, err = utilities.DecodeResults[data.GetIncrementResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetIncrementResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (s *SubscriptionsController) UpdateSubscriptionAffiliationId(
    subscriptionId string,
    request data.UpdateSubscriptionAffiliationIdRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetSubscriptionResponse],
    error) {
    req := s.prepareRequest(
      "PATCH",
      fmt.Sprintf("/subscriptions/%s/gateway-affiliation-id", subscriptionId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionResponse]{}, err
    }
    
    var result data.GetSubscriptionResponse
    result, err = utilities.DecodeResults[data.GetSubscriptionResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Updates the metadata from a subscription
func (s *SubscriptionsController) UpdateSubscriptionMetadata(
    subscriptionId string,
    request data.UpdateMetadataRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetSubscriptionResponse],
    error) {
    req := s.prepareRequest(
      "PATCH",
      fmt.Sprintf("/Subscriptions/%s/metadata", subscriptionId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionResponse]{}, err
    }
    
    var result data.GetSubscriptionResponse
    result, err = utilities.DecodeResults[data.GetSubscriptionResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Deletes a increment
func (s *SubscriptionsController) DeleteIncrement(
    subscriptionId string,
    incrementId string,
    idempotencyKey string) (
    https.ApiResponse[data.GetIncrementResponse],
    error) {
    req := s.prepareRequest(
      "DELETE",
      fmt.Sprintf("/subscriptions/%s/increments/%s", subscriptionId, incrementId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetIncrementResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetIncrementResponse]{}, err
    }
    
    var result data.GetIncrementResponse
    result, err = utilities.DecodeResults[data.GetIncrementResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetIncrementResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (s *SubscriptionsController) GetSubscriptionCycles(
    subscriptionId string,
    page string,
    size string) (
    https.ApiResponse[data.ListCyclesResponse],
    error) {
    req := s.prepareRequest(
      "GET",
      fmt.Sprintf("/subscriptions/%s/cycles", subscriptionId),
    )
    req.Authenticate(true)
    req.QueryParam("page", page)
    req.QueryParam("size", size)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.ListCyclesResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.ListCyclesResponse]{}, err
    }
    
    var result data.ListCyclesResponse
    result, err = utilities.DecodeResults[data.ListCyclesResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.ListCyclesResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (s *SubscriptionsController) GetDiscounts(
    subscriptionId string,
    page int,
    size int) (
    https.ApiResponse[data.ListDiscountsResponse],
    error) {
    req := s.prepareRequest(
      "GET",
      fmt.Sprintf("/subscriptions/%s/discounts/", subscriptionId),
    )
    req.Authenticate(true)
    req.QueryParam("page", page)
    req.QueryParam("size", size)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.ListDiscountsResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.ListDiscountsResponse]{}, err
    }
    
    var result data.ListDiscountsResponse
    result, err = utilities.DecodeResults[data.ListDiscountsResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.ListDiscountsResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Updates the billing date from a subscription
func (s *SubscriptionsController) UpdateSubscriptionBillingDate(
    subscriptionId string,
    request data.UpdateSubscriptionBillingDateRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetSubscriptionResponse],
    error) {
    req := s.prepareRequest(
      "PATCH",
      fmt.Sprintf("/subscriptions/%s/billing-date", subscriptionId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionResponse]{}, err
    }
    
    var result data.GetSubscriptionResponse
    result, err = utilities.DecodeResults[data.GetSubscriptionResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Deletes a subscription item
func (s *SubscriptionsController) DeleteSubscriptionItem(
    subscriptionId string,
    subscriptionItemId string,
    idempotencyKey string) (
    https.ApiResponse[data.GetSubscriptionItemResponse],
    error) {
    req := s.prepareRequest(
      "DELETE",
      fmt.Sprintf("/subscriptions/%s/items/%s", subscriptionId, subscriptionItemId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionItemResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionItemResponse]{}, err
    }
    
    var result data.GetSubscriptionItemResponse
    result, err = utilities.DecodeResults[data.GetSubscriptionItemResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionItemResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (s *SubscriptionsController) GetIncrements(
    subscriptionId string,
    page int,
    size int) (
    https.ApiResponse[data.ListIncrementsResponse],
    error) {
    req := s.prepareRequest(
      "GET",
      fmt.Sprintf("/subscriptions/%s/increments/", subscriptionId),
    )
    req.Authenticate(true)
    req.QueryParam("page", page)
    req.QueryParam("size", size)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.ListIncrementsResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.ListIncrementsResponse]{}, err
    }
    
    var result data.ListIncrementsResponse
    result, err = utilities.DecodeResults[data.ListIncrementsResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.ListIncrementsResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Updates the boleto due days from a subscription
func (s *SubscriptionsController) UpdateSubscriptionDueDays(
    subscriptionId string,
    request data.UpdateSubscriptionDueDaysRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetSubscriptionResponse],
    error) {
    req := s.prepareRequest(
      "PATCH",
      fmt.Sprintf("/subscriptions/%s/boleto-due-days", subscriptionId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionResponse]{}, err
    }
    
    var result data.GetSubscriptionResponse
    result, err = utilities.DecodeResults[data.GetSubscriptionResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Updates the start at date from a subscription
func (s *SubscriptionsController) UpdateSubscriptionStartAt(
    subscriptionId string,
    request data.UpdateSubscriptionStartAtRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetSubscriptionResponse],
    error) {
    req := s.prepareRequest(
      "PATCH",
      fmt.Sprintf("/subscriptions/%s/start-at", subscriptionId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionResponse]{}, err
    }
    
    var result data.GetSubscriptionResponse
    result, err = utilities.DecodeResults[data.GetSubscriptionResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Updates a subscription item
func (s *SubscriptionsController) UpdateSubscriptionItem(
    subscriptionId string,
    itemId string,
    body data.UpdateSubscriptionItemRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetSubscriptionItemResponse],
    error) {
    req := s.prepareRequest(
      "PUT",
      fmt.Sprintf("/subscriptions/%s/items/%s", subscriptionId, itemId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(body)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionItemResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionItemResponse]{}, err
    }
    
    var result data.GetSubscriptionItemResponse
    result, err = utilities.DecodeResults[data.GetSubscriptionItemResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionItemResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Creates a new Subscription item
func (s *SubscriptionsController) CreateSubscriptionItem(
    subscriptionId string,
    request data.CreateSubscriptionItemRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetSubscriptionItemResponse],
    error) {
    req := s.prepareRequest(
      "POST",
      fmt.Sprintf("/subscriptions/%s/items", subscriptionId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionItemResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionItemResponse]{}, err
    }
    
    var result data.GetSubscriptionItemResponse
    result, err = utilities.DecodeResults[data.GetSubscriptionItemResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionItemResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Gets a subscription
func (s *SubscriptionsController) GetSubscription(subscriptionId string) (
    https.ApiResponse[data.GetSubscriptionResponse],
    error) {
    req := s.prepareRequest(
      "GET",
      fmt.Sprintf("/subscriptions/%s", subscriptionId),
    )
    req.Authenticate(true)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionResponse]{}, err
    }
    
    var result data.GetSubscriptionResponse
    result, err = utilities.DecodeResults[data.GetSubscriptionResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Lists all usages from a subscription item
func (s *SubscriptionsController) GetUsages(
    subscriptionId string,
    itemId string,
    page int,
    size int,
    code string,
    group string,
    usedSince time.Time,
    usedUntil time.Time) (
    https.ApiResponse[data.ListUsagesResponse],
    error) {
    req := s.prepareRequest(
      "GET",
      fmt.Sprintf("/subscriptions/%s/items/%s/usages", subscriptionId, itemId),
    )
    req.Authenticate(true)
    req.QueryParam("page", page)
    req.QueryParam("size", size)
    req.QueryParam("code", code)
    req.QueryParam("group", group)
    req.QueryParam("used_since", usedSince.Format(time.RFC3339))
    req.QueryParam("used_until", usedUntil.Format(time.RFC3339))
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.ListUsagesResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.ListUsagesResponse]{}, err
    }
    
    var result data.ListUsagesResponse
    result, err = utilities.DecodeResults[data.ListUsagesResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.ListUsagesResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (s *SubscriptionsController) UpdateLatestPeriodEndAt(
    subscriptionId string,
    request data.UpdateCurrentCycleEndDateRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetSubscriptionResponse],
    error) {
    req := s.prepareRequest(
      "PATCH",
      fmt.Sprintf("/subscriptions/%s/periods/latest/end-at", subscriptionId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionResponse]{}, err
    }
    
    var result data.GetSubscriptionResponse
    result, err = utilities.DecodeResults[data.GetSubscriptionResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Atualização do valor mínimo da assinatura
func (s *SubscriptionsController) UpdateSubscriptionMiniumPrice(
    subscriptionId string,
    request data.UpdateSubscriptionMinimumPriceRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetSubscriptionResponse],
    error) {
    req := s.prepareRequest(
      "PATCH",
      fmt.Sprintf("/subscriptions/%s/minimum_price", subscriptionId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionResponse]{}, err
    }
    
    var result data.GetSubscriptionResponse
    result, err = utilities.DecodeResults[data.GetSubscriptionResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (s *SubscriptionsController) GetSubscriptionCycleById(
    subscriptionId string,
    cycleId string) (
    https.ApiResponse[data.GetPeriodResponse],
    error) {
    req := s.prepareRequest(
      "GET",
      fmt.Sprintf("/subscriptions/%s/cycles/%s", subscriptionId, cycleId),
    )
    req.Authenticate(true)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetPeriodResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetPeriodResponse]{}, err
    }
    
    var result data.GetPeriodResponse
    result, err = utilities.DecodeResults[data.GetPeriodResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetPeriodResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (s *SubscriptionsController) GetUsageReport(
    subscriptionId string,
    periodId string) (
    https.ApiResponse[data.GetUsageReportResponse],
    error) {
    req := s.prepareRequest(
      "GET",
      fmt.Sprintf("/subscriptions/%s/periods/%s/usages/report", subscriptionId, periodId),
    )
    req.Authenticate(true)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetUsageReportResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetUsageReportResponse]{}, err
    }
    
    var result data.GetUsageReportResponse
    result, err = utilities.DecodeResults[data.GetUsageReportResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetUsageReportResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (s *SubscriptionsController) UpdateSplitSubscription(
    id string,
    request data.UpdateSubscriptionSplitRequest) (
    https.ApiResponse[data.GetSubscriptionResponse],
    error) {
    req := s.prepareRequest("PATCH", fmt.Sprintf("/subscriptions/%s/split", id))
    req.Authenticate(true)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionResponse]{}, err
    }
    
    var result data.GetSubscriptionResponse
    result, err = utilities.DecodeResults[data.GetSubscriptionResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetSubscriptionResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}
