package controllers

import (
    "fmt"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "pagarmeapisdk/data"
)

type TransactionsController struct {
    baseController
}

func NewTransactionsController(baseController baseController) *TransactionsController {
    transactionsController := TransactionsController{baseController: baseController}
    return &transactionsController
}

// TODO: type endpoint description here
func (t *TransactionsController) GetTransaction(transactionId string) (
    https.ApiResponse[data.GetTransactionResponseInterface],
    error) {
    req := t.prepareRequest("GET", fmt.Sprintf("/transactions/%s", transactionId))
    req.Authenticate(true)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetTransactionResponseInterface]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetTransactionResponseInterface]{}, err
    }
    
    var result data.GetTransactionResponseInterface
    result, err = utilities.DecodeResults[*data.GetTransactionResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetTransactionResponseInterface]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}
