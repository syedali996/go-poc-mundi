package controllers

import (
    "fmt"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "pagarmeapisdk/data"
    "time"
)

type OrdersController struct {
    baseController
}

func NewOrdersController(baseController baseController) *OrdersController {
    ordersController := OrdersController{baseController: baseController}
    return &ordersController
}

// Gets all orders
func (o *OrdersController) GetOrders(
    page int,
    size int,
    code string,
    status string,
    createdSince time.Time,
    createdUntil time.Time,
    customerId string) (
    https.ApiResponse[data.ListOrderResponse],
    error) {
    req := o.prepareRequest("GET", "/orders")
    req.Authenticate(true)
    req.QueryParam("page", page)
    req.QueryParam("size", size)
    req.QueryParam("code", code)
    req.QueryParam("status", status)
    req.QueryParam("created_since", createdSince.Format(time.RFC3339))
    req.QueryParam("created_until", createdUntil.Format(time.RFC3339))
    req.QueryParam("customer_id", customerId)
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.ListOrderResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.ListOrderResponse]{}, err
    }
    
    var result data.ListOrderResponse
    result, err = utilities.DecodeResults[data.ListOrderResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.ListOrderResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (o *OrdersController) UpdateOrderItem(
    orderId string,
    itemId string,
    request data.UpdateOrderItemRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetOrderItemResponse],
    error) {
    req := o.prepareRequest(
      "PUT",
      fmt.Sprintf("/orders/%s/items/%s", orderId, itemId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetOrderItemResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetOrderItemResponse]{}, err
    }
    
    var result data.GetOrderItemResponse
    result, err = utilities.DecodeResults[data.GetOrderItemResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetOrderItemResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (o *OrdersController) DeleteAllOrderItems(
    orderId string,
    idempotencyKey string) (
    https.ApiResponse[data.GetOrderResponse],
    error) {
    req := o.prepareRequest("DELETE", fmt.Sprintf("/orders/%s/items", orderId))
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetOrderResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetOrderResponse]{}, err
    }
    
    var result data.GetOrderResponse
    result, err = utilities.DecodeResults[data.GetOrderResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetOrderResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (o *OrdersController) DeleteOrderItem(
    orderId string,
    itemId string,
    idempotencyKey string) (
    https.ApiResponse[data.GetOrderItemResponse],
    error) {
    req := o.prepareRequest(
      "DELETE",
      fmt.Sprintf("/orders/%s/items/%s", orderId, itemId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetOrderItemResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetOrderItemResponse]{}, err
    }
    
    var result data.GetOrderItemResponse
    result, err = utilities.DecodeResults[data.GetOrderItemResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetOrderItemResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (o *OrdersController) CloseOrder(
    id string,
    request data.UpdateOrderStatusRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetOrderResponse],
    error) {
    req := o.prepareRequest("PATCH", fmt.Sprintf("/orders/%s/closed", id))
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetOrderResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetOrderResponse]{}, err
    }
    
    var result data.GetOrderResponse
    result, err = utilities.DecodeResults[data.GetOrderResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetOrderResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Creates a new Order
func (o *OrdersController) CreateOrder(
    body data.CreateOrderRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetOrderResponse],
    error) {
    req := o.prepareRequest("POST", "/orders")
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(body)
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetOrderResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetOrderResponse]{}, err
    }
    
    var result data.GetOrderResponse
    result, err = utilities.DecodeResults[data.GetOrderResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetOrderResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (o *OrdersController) CreateOrderItem(
    orderId string,
    request data.CreateOrderItemRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetOrderItemResponse],
    error) {
    req := o.prepareRequest("POST", fmt.Sprintf("/orders/%s/items", orderId))
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetOrderItemResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetOrderItemResponse]{}, err
    }
    
    var result data.GetOrderItemResponse
    result, err = utilities.DecodeResults[data.GetOrderItemResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetOrderItemResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (o *OrdersController) GetOrderItem(
    orderId string,
    itemId string) (
    https.ApiResponse[data.GetOrderItemResponse],
    error) {
    req := o.prepareRequest(
      "GET",
      fmt.Sprintf("/orders/%s/items/%s", orderId, itemId),
    )
    req.Authenticate(true)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetOrderItemResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetOrderItemResponse]{}, err
    }
    
    var result data.GetOrderItemResponse
    result, err = utilities.DecodeResults[data.GetOrderItemResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetOrderItemResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Updates the metadata from an order
func (o *OrdersController) UpdateOrderMetadata(
    orderId string,
    request data.UpdateMetadataRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetOrderResponse],
    error) {
    req := o.prepareRequest("PATCH", fmt.Sprintf("/Orders/%s/metadata", orderId))
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetOrderResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetOrderResponse]{}, err
    }
    
    var result data.GetOrderResponse
    result, err = utilities.DecodeResults[data.GetOrderResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetOrderResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Gets an order
func (o *OrdersController) GetOrder(orderId string) (
    https.ApiResponse[data.GetOrderResponse],
    error) {
    req := o.prepareRequest("GET", fmt.Sprintf("/orders/%s", orderId))
    req.Authenticate(true)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetOrderResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetOrderResponse]{}, err
    }
    
    var result data.GetOrderResponse
    result, err = utilities.DecodeResults[data.GetOrderResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetOrderResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}
