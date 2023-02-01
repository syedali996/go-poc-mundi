package controllers

import (
    "fmt"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "pagarmeapisdk/data"
)

type TransfersController struct {
    baseController
}

func NewTransfersController(baseController baseController) *TransfersController {
    transfersController := TransfersController{baseController: baseController}
    return &transfersController
}

// TODO: type endpoint description here
func (t *TransfersController) GetTransferById(transferId string) (
    https.ApiResponse[data.GetTransfer],
    error) {
    req := t.prepareRequest("GET", fmt.Sprintf("/transfers/%s", transferId))
    req.Authenticate(true)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetTransfer]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetTransfer]{}, err
    }
    
    var result data.GetTransfer
    result, err = utilities.DecodeResults[data.GetTransfer](decoder)
    if err != nil {
        return https.ApiResponse[data.GetTransfer]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (t *TransfersController) CreateTransfer(request data.CreateTransfer) (
    https.ApiResponse[data.GetTransfer],
    error) {
    req := t.prepareRequest("POST", "/transfers/recipients")
    req.Authenticate(true)
    req.Json(request)
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetTransfer]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetTransfer]{}, err
    }
    
    var result data.GetTransfer
    result, err = utilities.DecodeResults[data.GetTransfer](decoder)
    if err != nil {
        return https.ApiResponse[data.GetTransfer]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Gets all transfers
func (t *TransfersController) GetTransfers() (
    https.ApiResponse[data.ListTransfers],
    error) {
    req := t.prepareRequest("GET", "/transfers")
    req.Authenticate(true)
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.ListTransfers]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.ListTransfers]{}, err
    }
    
    var result data.ListTransfers
    result, err = utilities.DecodeResults[data.ListTransfers](decoder)
    if err != nil {
        return https.ApiResponse[data.ListTransfers]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}
