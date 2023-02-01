package controllers

import (
    "fmt"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "pagarmeapisdk/data"
    "time"
)

type PlansController struct {
    baseController
}

func NewPlansController(baseController baseController) *PlansController {
    plansController := PlansController{baseController: baseController}
    return &plansController
}

// Gets a plan
func (p *PlansController) GetPlan(planId string) (
    https.ApiResponse[data.GetPlanResponse],
    error) {
    req := p.prepareRequest("GET", fmt.Sprintf("/plans/%s", planId))
    req.Authenticate(true)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetPlanResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetPlanResponse]{}, err
    }
    
    var result data.GetPlanResponse
    result, err = utilities.DecodeResults[data.GetPlanResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetPlanResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Deletes a plan
func (p *PlansController) DeletePlan(
    planId string,
    idempotencyKey string) (
    https.ApiResponse[data.GetPlanResponse],
    error) {
    req := p.prepareRequest("DELETE", fmt.Sprintf("/plans/%s", planId))
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetPlanResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetPlanResponse]{}, err
    }
    
    var result data.GetPlanResponse
    result, err = utilities.DecodeResults[data.GetPlanResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetPlanResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Updates the metadata from a plan
func (p *PlansController) UpdatePlanMetadata(
    planId string,
    request data.UpdateMetadataRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetPlanResponse],
    error) {
    req := p.prepareRequest("PATCH", fmt.Sprintf("/Plans/%s/metadata", planId))
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetPlanResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetPlanResponse]{}, err
    }
    
    var result data.GetPlanResponse
    result, err = utilities.DecodeResults[data.GetPlanResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetPlanResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Updates a plan item
func (p *PlansController) UpdatePlanItem(
    planId string,
    planItemId string,
    body data.UpdatePlanItemRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetPlanItemResponse],
    error) {
    req := p.prepareRequest(
      "PUT",
      fmt.Sprintf("/plans/%s/items/%s", planId, planItemId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(body)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetPlanItemResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetPlanItemResponse]{}, err
    }
    
    var result data.GetPlanItemResponse
    result, err = utilities.DecodeResults[data.GetPlanItemResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetPlanItemResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Adds a new item to a plan
func (p *PlansController) CreatePlanItem(
    planId string,
    request data.CreatePlanItemRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetPlanItemResponse],
    error) {
    req := p.prepareRequest("POST", fmt.Sprintf("/plans/%s/items", planId))
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetPlanItemResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetPlanItemResponse]{}, err
    }
    
    var result data.GetPlanItemResponse
    result, err = utilities.DecodeResults[data.GetPlanItemResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetPlanItemResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Gets a plan item
func (p *PlansController) GetPlanItem(
    planId string,
    planItemId string) (
    https.ApiResponse[data.GetPlanItemResponse],
    error) {
    req := p.prepareRequest(
      "GET",
      fmt.Sprintf("/plans/%s/items/%s", planId, planItemId),
    )
    req.Authenticate(true)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetPlanItemResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetPlanItemResponse]{}, err
    }
    
    var result data.GetPlanItemResponse
    result, err = utilities.DecodeResults[data.GetPlanItemResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetPlanItemResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Creates a new plan
func (p *PlansController) CreatePlan(
    body data.CreatePlanRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetPlanResponse],
    error) {
    req := p.prepareRequest("POST", "/plans")
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(body)
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetPlanResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetPlanResponse]{}, err
    }
    
    var result data.GetPlanResponse
    result, err = utilities.DecodeResults[data.GetPlanResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetPlanResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Removes an item from a plan
func (p *PlansController) DeletePlanItem(
    planId string,
    planItemId string,
    idempotencyKey string) (
    https.ApiResponse[data.GetPlanItemResponse],
    error) {
    req := p.prepareRequest(
      "DELETE",
      fmt.Sprintf("/plans/%s/items/%s", planId, planItemId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetPlanItemResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetPlanItemResponse]{}, err
    }
    
    var result data.GetPlanItemResponse
    result, err = utilities.DecodeResults[data.GetPlanItemResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetPlanItemResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Gets all plans
func (p *PlansController) GetPlans(
    page int,
    size int,
    name string,
    status string,
    billingType string,
    createdSince time.Time,
    createdUntil time.Time) (
    https.ApiResponse[data.ListPlansResponse],
    error) {
    req := p.prepareRequest("GET", "/plans")
    req.Authenticate(true)
    req.QueryParam("page", page)
    req.QueryParam("size", size)
    req.QueryParam("name", name)
    req.QueryParam("status", status)
    req.QueryParam("billing_type", billingType)
    req.QueryParam("created_since", createdSince.Format(time.RFC3339))
    req.QueryParam("created_until", createdUntil.Format(time.RFC3339))
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.ListPlansResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.ListPlansResponse]{}, err
    }
    
    var result data.ListPlansResponse
    result, err = utilities.DecodeResults[data.ListPlansResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.ListPlansResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Updates a plan
func (p *PlansController) UpdatePlan(
    planId string,
    request data.UpdatePlanRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetPlanResponse],
    error) {
    req := p.prepareRequest("PUT", fmt.Sprintf("/plans/%s", planId))
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetPlanResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetPlanResponse]{}, err
    }
    
    var result data.GetPlanResponse
    result, err = utilities.DecodeResults[data.GetPlanResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetPlanResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}
