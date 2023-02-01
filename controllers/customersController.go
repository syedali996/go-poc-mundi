package controllers

import (
    "fmt"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "pagarmeapisdk/data"
)

type CustomersController struct {
    baseController
}

func NewCustomersController(baseController baseController) *CustomersController {
    customersController := CustomersController{baseController: baseController}
    return &customersController
}

// Updates a card
func (c *CustomersController) UpdateCard(
    customerId string,
    cardId string,
    request data.UpdateCardRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetCardResponse],
    error) {
    req := c.prepareRequest(
      "PUT",
      fmt.Sprintf("/customers/%s/cards/%s", customerId, cardId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetCardResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetCardResponse]{}, err
    }
    
    var result data.GetCardResponse
    result, err = utilities.DecodeResults[data.GetCardResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetCardResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Updates an address
func (c *CustomersController) UpdateAddress(
    customerId string,
    addressId string,
    request data.UpdateAddressRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetAddressResponse],
    error) {
    req := c.prepareRequest(
      "PUT",
      fmt.Sprintf("/customers/%s/addresses/%s", customerId, addressId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetAddressResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetAddressResponse]{}, err
    }
    
    var result data.GetAddressResponse
    result, err = utilities.DecodeResults[data.GetAddressResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetAddressResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Delete a customer's access token
func (c *CustomersController) DeleteAccessToken(
    customerId string,
    tokenId string,
    idempotencyKey string) (
    https.ApiResponse[data.GetAccessTokenResponse],
    error) {
    req := c.prepareRequest(
      "DELETE",
      fmt.Sprintf("/customers/%s/access-tokens/%s", customerId, tokenId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetAccessTokenResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetAccessTokenResponse]{}, err
    }
    
    var result data.GetAccessTokenResponse
    result, err = utilities.DecodeResults[data.GetAccessTokenResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetAccessTokenResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Creates a new customer
func (c *CustomersController) CreateCustomer(
    request data.CreateCustomerRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetCustomerResponse],
    error) {
    req := c.prepareRequest("POST", "/customers")
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetCustomerResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetCustomerResponse]{}, err
    }
    
    var result data.GetCustomerResponse
    result, err = utilities.DecodeResults[data.GetCustomerResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetCustomerResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Creates a new address for a customer
func (c *CustomersController) CreateAddress(
    customerId string,
    request data.CreateAddressRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetAddressResponse],
    error) {
    req := c.prepareRequest(
      "POST",
      fmt.Sprintf("/customers/%s/addresses", customerId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetAddressResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetAddressResponse]{}, err
    }
    
    var result data.GetAddressResponse
    result, err = utilities.DecodeResults[data.GetAddressResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetAddressResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Delete a Customer's access tokens
func (c *CustomersController) DeleteAccessTokens(customerId string) (
    https.ApiResponse[data.ListAccessTokensResponse],
    error) {
    req := c.prepareRequest(
      "GET",
      fmt.Sprintf("/customers/%s/access-tokens/", customerId),
    )
    req.Authenticate(true)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.ListAccessTokensResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.ListAccessTokensResponse]{}, err
    }
    
    var result data.ListAccessTokensResponse
    result, err = utilities.DecodeResults[data.ListAccessTokensResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.ListAccessTokensResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Get a customer's address
func (c *CustomersController) GetAddress(
    customerId string,
    addressId string) (
    https.ApiResponse[data.GetAddressResponse],
    error) {
    req := c.prepareRequest(
      "GET",
      fmt.Sprintf("/customers/%s/addresses/%s", customerId, addressId),
    )
    req.Authenticate(true)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetAddressResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetAddressResponse]{}, err
    }
    
    var result data.GetAddressResponse
    result, err = utilities.DecodeResults[data.GetAddressResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetAddressResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Delete a Customer's address
func (c *CustomersController) DeleteAddress(
    customerId string,
    addressId string,
    idempotencyKey string) (
    https.ApiResponse[data.GetAddressResponse],
    error) {
    req := c.prepareRequest(
      "DELETE",
      fmt.Sprintf("/customers/%s/addresses/%s", customerId, addressId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetAddressResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetAddressResponse]{}, err
    }
    
    var result data.GetAddressResponse
    result, err = utilities.DecodeResults[data.GetAddressResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetAddressResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Creates a new card for a customer
func (c *CustomersController) CreateCard(
    customerId string,
    request data.CreateCardRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetCardResponse],
    error) {
    req := c.prepareRequest("POST", fmt.Sprintf("/customers/%s/cards", customerId))
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetCardResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetCardResponse]{}, err
    }
    
    var result data.GetCardResponse
    result, err = utilities.DecodeResults[data.GetCardResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetCardResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Get all Customers
func (c *CustomersController) GetCustomers(
    name string,
    document string,
    page int,
    size int,
    email string,
    code string) (
    https.ApiResponse[data.ListCustomersResponse],
    error) {
    req := c.prepareRequest("GET", "/customers")
    req.Authenticate(true)
    req.QueryParam("name", name)
    req.QueryParam("document", document)
    req.QueryParam("page", page)
    req.QueryParam("size", size)
    req.QueryParam("email", email)
    req.QueryParam("Code", code)
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.ListCustomersResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.ListCustomersResponse]{}, err
    }
    
    var result data.ListCustomersResponse
    result, err = utilities.DecodeResults[data.ListCustomersResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.ListCustomersResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Updates a customer
func (c *CustomersController) UpdateCustomer(
    customerId string,
    request data.UpdateCustomerRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetCustomerResponse],
    error) {
    req := c.prepareRequest("PUT", fmt.Sprintf("/customers/%s", customerId))
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetCustomerResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetCustomerResponse]{}, err
    }
    
    var result data.GetCustomerResponse
    result, err = utilities.DecodeResults[data.GetCustomerResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetCustomerResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Creates a access token for a customer
func (c *CustomersController) CreateAccessToken(
    customerId string,
    request data.CreateAccessTokenRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetAccessTokenResponse],
    error) {
    req := c.prepareRequest(
      "POST",
      fmt.Sprintf("/customers/%s/access-tokens", customerId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetAccessTokenResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetAccessTokenResponse]{}, err
    }
    
    var result data.GetAccessTokenResponse
    result, err = utilities.DecodeResults[data.GetAccessTokenResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetAccessTokenResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Get all access tokens from a customer
func (c *CustomersController) GetAccessTokens(
    customerId string,
    page int,
    size int) (
    https.ApiResponse[data.ListAccessTokensResponse],
    error) {
    req := c.prepareRequest(
      "GET",
      fmt.Sprintf("/customers/%s/access-tokens", customerId),
    )
    req.Authenticate(true)
    req.QueryParam("page", page)
    req.QueryParam("size", size)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.ListAccessTokensResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.ListAccessTokensResponse]{}, err
    }
    
    var result data.ListAccessTokensResponse
    result, err = utilities.DecodeResults[data.ListAccessTokensResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.ListAccessTokensResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Get all cards from a customer
func (c *CustomersController) GetCards(
    customerId string,
    page int,
    size int) (
    https.ApiResponse[data.ListCardsResponse],
    error) {
    req := c.prepareRequest("GET", fmt.Sprintf("/customers/%s/cards", customerId))
    req.Authenticate(true)
    req.QueryParam("page", page)
    req.QueryParam("size", size)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.ListCardsResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.ListCardsResponse]{}, err
    }
    
    var result data.ListCardsResponse
    result, err = utilities.DecodeResults[data.ListCardsResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.ListCardsResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Renew a card
func (c *CustomersController) RenewCard(
    customerId string,
    cardId string,
    idempotencyKey string) (
    https.ApiResponse[data.GetCardResponse],
    error) {
    req := c.prepareRequest(
      "POST",
      fmt.Sprintf("/customers/%s/cards/%s/renew", customerId, cardId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetCardResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetCardResponse]{}, err
    }
    
    var result data.GetCardResponse
    result, err = utilities.DecodeResults[data.GetCardResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetCardResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Get a Customer's access token
func (c *CustomersController) GetAccessToken(
    customerId string,
    tokenId string) (
    https.ApiResponse[data.GetAccessTokenResponse],
    error) {
    req := c.prepareRequest(
      "GET",
      fmt.Sprintf("/customers/%s/access-tokens/%s", customerId, tokenId),
    )
    req.Authenticate(true)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetAccessTokenResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetAccessTokenResponse]{}, err
    }
    
    var result data.GetAccessTokenResponse
    result, err = utilities.DecodeResults[data.GetAccessTokenResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetAccessTokenResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Updates the metadata a customer
func (c *CustomersController) UpdateCustomerMetadata(
    customerId string,
    request data.UpdateMetadataRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetCustomerResponse],
    error) {
    req := c.prepareRequest(
      "PATCH",
      fmt.Sprintf("/Customers/%s/metadata", customerId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    req.Json(request)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetCustomerResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetCustomerResponse]{}, err
    }
    
    var result data.GetCustomerResponse
    result, err = utilities.DecodeResults[data.GetCustomerResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetCustomerResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Delete a customer's card
func (c *CustomersController) DeleteCard(
    customerId string,
    cardId string,
    idempotencyKey string) (
    https.ApiResponse[data.GetCardResponse],
    error) {
    req := c.prepareRequest(
      "DELETE",
      fmt.Sprintf("/customers/%s/cards/%s", customerId, cardId),
    )
    req.Authenticate(true)
    req.Header("idempotency-key", idempotencyKey)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetCardResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetCardResponse]{}, err
    }
    
    var result data.GetCardResponse
    result, err = utilities.DecodeResults[data.GetCardResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetCardResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Gets all adressess from a customer
func (c *CustomersController) GetAddresses(
    customerId string,
    page int,
    size int) (
    https.ApiResponse[data.ListAddressesResponse],
    error) {
    req := c.prepareRequest(
      "GET",
      fmt.Sprintf("/customers/%s/addresses", customerId),
    )
    req.Authenticate(true)
    req.QueryParam("page", page)
    req.QueryParam("size", size)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.ListAddressesResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.ListAddressesResponse]{}, err
    }
    
    var result data.ListAddressesResponse
    result, err = utilities.DecodeResults[data.ListAddressesResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.ListAddressesResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Get a customer
func (c *CustomersController) GetCustomer(customerId string) (
    https.ApiResponse[data.GetCustomerResponse],
    error) {
    req := c.prepareRequest("GET", fmt.Sprintf("/customers/%s", customerId))
    req.Authenticate(true)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetCustomerResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetCustomerResponse]{}, err
    }
    
    var result data.GetCustomerResponse
    result, err = utilities.DecodeResults[data.GetCustomerResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetCustomerResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Get a customer's card
func (c *CustomersController) GetCard(
    customerId string,
    cardId string) (
    https.ApiResponse[data.GetCardResponse],
    error) {
    req := c.prepareRequest(
      "GET",
      fmt.Sprintf("/customers/%s/cards/%s", customerId, cardId),
    )
    req.Authenticate(true)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[data.GetCardResponse]{}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[data.GetCardResponse]{}, err
    }
    
    var result data.GetCardResponse
    result, err = utilities.DecodeResults[data.GetCardResponse](decoder)
    if err != nil {
        return https.ApiResponse[data.GetCardResponse]{}, err
    }
    
    return https.NewApiResponse(result, resp), err
}
