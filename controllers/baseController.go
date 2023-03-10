/*
Package pagarmeapisdk

This file was automatically generated by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package controllers

import (
    "github.com/apimatic/go-core-runtime/https"
    "net/http"
    "pagarmeapisdk/exceptions"
)

type callBuilderFactory interface {
	GetCallBuilder() https.CallBuilderFactory
}

type baseController struct {
	callBuilder    callBuilderFactory
	prepareRequest https.CallBuilderFactory
}

func NewBaseController(cb callBuilderFactory) *baseController {
	baseController := baseController{callBuilder: cb}
	baseController.prepareRequest = baseController.callBuilder.GetCallBuilder()
	return &baseController
}

func validateResponse(response http.Response) error {
	if response.StatusCode == 400 {
		return exceptions.NewError(400, "Invalid request")
	}
	if response.StatusCode == 401 {
		return exceptions.NewError(401, "Invalid API key")
	}
	if response.StatusCode == 404 {
		return exceptions.NewError(404, "An informed resource was not found")
	}
	if response.StatusCode == 412 {
		return exceptions.NewError(412, "Business validation error")
	}
	if response.StatusCode == 422 {
		return exceptions.NewError(422, "Contract validation error")
	}
	if response.StatusCode == 500 {
		return exceptions.NewError(500, "Internal server error")
	}
	return nil
}
