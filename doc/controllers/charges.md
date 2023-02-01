# Charges

```go
chargesController := client.ChargesController
```

## Class Name

`ChargesController`

## Methods

* [Update Charge Metadata](../../doc/controllers/charges.md#update-charge-metadata)
* [Update Charge Payment Method](../../doc/controllers/charges.md#update-charge-payment-method)
* [Get Charge Transactions](../../doc/controllers/charges.md#get-charge-transactions)
* [Update Charge Due Date](../../doc/controllers/charges.md#update-charge-due-date)
* [Get Charges](../../doc/controllers/charges.md#get-charges)
* [Capture Charge](../../doc/controllers/charges.md#capture-charge)
* [Update Charge Card](../../doc/controllers/charges.md#update-charge-card)
* [Get Charge](../../doc/controllers/charges.md#get-charge)
* [Get Charges Summary](../../doc/controllers/charges.md#get-charges-summary)
* [Retry Charge](../../doc/controllers/charges.md#retry-charge)
* [Cancel Charge](../../doc/controllers/charges.md#cancel-charge)
* [Create Charge](../../doc/controllers/charges.md#create-charge)
* [Confirm Payment](../../doc/controllers/charges.md#confirm-payment)


# Update Charge Metadata

Updates the metadata from a charge

```go
UpdateChargeMetadata(
    chargeId string,
    request data.UpdateMetadataRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetChargeResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `chargeId` | `string` | Template, Required | The charge id |
| `request` | [`data.UpdateMetadataRequest`](../../doc/models/update-metadata-request.md) | Body, Required | Request for updating the charge metadata |
| `idempotencyKey` | `string` | Header, Optional | - |

## Response Type

[`data.GetChargeResponse`](../../doc/models/get-charge-response.md)

## Example Usage

```go
chargeId := "charge_id8"
requestMetadata := map[string]string{"Liquid error: Object of type 'System.Collections.Generic.KeyValuePair`2[System.String,System.Object]' cannot be converted to type 'System.String'."}
request := data.UpdateMetadataRequest { 
    Metadata: requestMetadata,
}

apiResponse, err := chargesController.UpdateChargeMetadata(chargeId, request, idempotencyKey)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Charge Payment Method

Updates a charge's payment method

```go
UpdateChargePaymentMethod(
    chargeId string,
    request data.UpdateChargePaymentMethodRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetChargeResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `chargeId` | `string` | Template, Required | Charge id |
| `request` | [`data.UpdateChargePaymentMethodRequest`](../../doc/models/update-charge-payment-method-request.md) | Body, Required | Request for updating the payment method from a charge |
| `idempotencyKey` | `string` | Header, Optional | - |

## Response Type

[`data.GetChargeResponse`](../../doc/models/get-charge-response.md)

## Example Usage

```go
chargeId := "charge_id8"
requestCreditCard := data.CreateCreditCardPaymentRequest { }

requestDebitCard := data.CreateDebitCardPaymentRequest { }

requestBoletoBillingAddressMetadata := map[string]string{"Liquid error: Object of type 'System.Collections.Generic.KeyValuePair`2[System.String,System.Object]' cannot be converted to type 'System.String'."}
requestBoletoBillingAddress := data.CreateAddressRequest { 
    Street: "street8",
    Number: "number4",
    ZipCode: "zip_code2",
    Neighborhood: "neighborhood4",
    City: "city2",
    State: "state6",
    Country: "country2",
    Complement: "complement6",
    Metadata: requestBoletoBillingAddressMetadata,
    Line1: "line_18",
    Line2: "line_26",
}

requestBoleto := data.CreateBoletoPaymentRequest { 
    Retries: 10,
    Bank: "bank4",
    Instructions: "instructions4",
    BillingAddress: requestBoletoBillingAddress,
    BillingAddressId: "billing_address_id2",
    DocumentNumber: "document_number0",
    StatementDescriptor: "statement_descriptor6",
}

requestVoucher := data.CreateVoucherPaymentRequest { }

requestCash := data.CreateCashPaymentRequest { 
    Description: "description6",
    Confirm: false,
}

requestBankTransfer := data.CreateBankTransferPaymentRequest { 
    Bank: "bank4",
    Retries: 204,
}

requestPrivateLabel := data.CreatePrivateLabelPaymentRequest { }

request := data.UpdateChargePaymentMethodRequest { 
    UpdateSubscription: false,
    PaymentMethod: "payment_method4",
    CreditCard: requestCreditCard,
    DebitCard: requestDebitCard,
    Boleto: requestBoleto,
    Voucher: requestVoucher,
    Cash: requestCash,
    BankTransfer: requestBankTransfer,
    PrivateLabel: requestPrivateLabel,
}

apiResponse, err := chargesController.UpdateChargePaymentMethod(chargeId, request, idempotencyKey)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Charge Transactions

```go
GetChargeTransactions(
    chargeId string,
    page int,
    size int) (
    https.ApiResponse[data.ListChargeTransactionsResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `chargeId` | `string` | Template, Required | Charge Id |
| `page` | `int` | Query, Optional | Page number |
| `size` | `int` | Query, Optional | Page size |

## Response Type

[`data.ListChargeTransactionsResponse`](../../doc/models/list-charge-transactions-response.md)

## Example Usage

```go
chargeId := "charge_id8"
apiResponse, err := chargesController.GetChargeTransactions(chargeId, page, size)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Charge Due Date

Updates the due date from a charge

```go
UpdateChargeDueDate(
    chargeId string,
    request data.UpdateChargeDueDateRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetChargeResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `chargeId` | `string` | Template, Required | Charge Id |
| `request` | [`data.UpdateChargeDueDateRequest`](../../doc/models/update-charge-due-date-request.md) | Body, Required | Request for updating the due date |
| `idempotencyKey` | `string` | Header, Optional | - |

## Response Type

[`data.GetChargeResponse`](../../doc/models/get-charge-response.md)

## Example Usage

```go
chargeId := "charge_id8"
request := data.UpdateChargeDueDateRequest { }

apiResponse, err := chargesController.UpdateChargeDueDate(chargeId, request, idempotencyKey)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Charges

Lists all charges

```go
GetCharges(
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
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `page` | `int` | Query, Optional | Page number |
| `size` | `int` | Query, Optional | Page size |
| `code` | `string` | Query, Optional | Filter for charge's code |
| `status` | `string` | Query, Optional | Filter for charge's status |
| `paymentMethod` | `string` | Query, Optional | Filter for charge's payment method |
| `customerId` | `string` | Query, Optional | Filter for charge's customer id |
| `orderId` | `string` | Query, Optional | Filter for charge's order id |
| `createdSince` | `time.Time` | Query, Optional | Filter for the beginning of the range for charge's creation |
| `createdUntil` | `time.Time` | Query, Optional | Filter for the end of the range for charge's creation |

## Response Type

[`data.ListChargesResponse`](../../doc/models/list-charges-response.md)

## Example Usage

```go
apiResponse, err := chargesController.GetCharges(page, size, code, status, paymentMethod, customerId, orderId, createdSince, createdUntil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Capture Charge

Captures a charge

```go
CaptureCharge(
    chargeId string,
    request data.CreateCaptureChargeRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetChargeResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `chargeId` | `string` | Template, Required | Charge id |
| `request` | [`data.CreateCaptureChargeRequest`](../../doc/models/create-capture-charge-request.md) | Body, Optional | Request for capturing a charge |
| `idempotencyKey` | `string` | Header, Optional | - |

## Response Type

[`data.GetChargeResponse`](../../doc/models/get-charge-response.md)

## Example Usage

```go
chargeId := "charge_id8"
apiResponse, err := chargesController.CaptureCharge(chargeId, request, idempotencyKey)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Charge Card

Updates the card from a charge

```go
UpdateChargeCard(
    chargeId string,
    request data.UpdateChargeCardRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetChargeResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `chargeId` | `string` | Template, Required | Charge id |
| `request` | [`data.UpdateChargeCardRequest`](../../doc/models/update-charge-card-request.md) | Body, Required | Request for updating a charge's card |
| `idempotencyKey` | `string` | Header, Optional | - |

## Response Type

[`data.GetChargeResponse`](../../doc/models/get-charge-response.md)

## Example Usage

```go
chargeId := "charge_id8"
requestCardBillingAddressMetadata := map[string]string{"Liquid error: Object of type 'System.Collections.Generic.KeyValuePair`2[System.String,System.Object]' cannot be converted to type 'System.String'."}
requestCardBillingAddress := data.CreateAddressRequest { 
    Street: "street2",
    Number: "number0",
    ZipCode: "zip_code6",
    Neighborhood: "neighborhood8",
    City: "city8",
    State: "state2",
    Country: "country6",
    Complement: "complement2",
    Metadata: requestCardBillingAddressMetadata,
    Line1: "line_14",
    Line2: "line_20",
}

requestCardMetadata := map[string]string{"Liquid error: Object of type 'System.Collections.Generic.KeyValuePair`2[System.String,System.Object]' cannot be converted to type 'System.String'.", "Liquid error: Object of type 'System.Collections.Generic.KeyValuePair`2[System.String,System.Object]' cannot be converted to type 'System.String'.", "Liquid error: Object of type 'System.Collections.Generic.KeyValuePair`2[System.String,System.Object]' cannot be converted to type 'System.String'."}
requestCardOptions := data.CreateCardOptionsRequest { 
    VerifyCard: false,
}

requestCard := data.CreateCardRequest { 
    Number: "number2",
    HolderName: "holder_name6",
    ExpMonth: 80,
    ExpYear: 216,
    Cvv: "cvv8",
    BillingAddress: requestCardBillingAddress,
    Brand: "brand4",
    BillingAddressId: "billing_address_id6",
    Metadata: requestCardMetadata,
    Type: "credit",
    Options: requestCardOptions,
    PrivateLabel: false,
    Label: "label0",
}

request := data.UpdateChargeCardRequest { 
    UpdateSubscription: false,
    CardId: "card_id2",
    Card: requestCard,
    Recurrence: false,
}

apiResponse, err := chargesController.UpdateChargeCard(chargeId, request, idempotencyKey)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Charge

Get a charge from its id

```go
GetCharge(
    chargeId string) (
    https.ApiResponse[data.GetChargeResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `chargeId` | `string` | Template, Required | Charge id |

## Response Type

[`data.GetChargeResponse`](../../doc/models/get-charge-response.md)

## Example Usage

```go
chargeId := "charge_id8"
apiResponse, err := chargesController.GetCharge(chargeId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Charges Summary

```go
GetChargesSummary(
    status string,
    createdSince time.Time,
    createdUntil time.Time) (
    https.ApiResponse[data.GetChargesSummaryResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `status` | `string` | Query, Required | - |
| `createdSince` | `time.Time` | Query, Optional | - |
| `createdUntil` | `time.Time` | Query, Optional | - |

## Response Type

[`data.GetChargesSummaryResponse`](../../doc/models/get-charges-summary-response.md)

## Example Usage

```go
status := "status8"
apiResponse, err := chargesController.GetChargesSummary(status, createdSince, createdUntil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Retry Charge

Retries a charge

```go
RetryCharge(
    chargeId string,
    idempotencyKey string) (
    https.ApiResponse[data.GetChargeResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `chargeId` | `string` | Template, Required | Charge id |
| `idempotencyKey` | `string` | Header, Optional | - |

## Response Type

[`data.GetChargeResponse`](../../doc/models/get-charge-response.md)

## Example Usage

```go
chargeId := "charge_id8"
apiResponse, err := chargesController.RetryCharge(chargeId, idempotencyKey)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Cancel Charge

Cancel a charge

```go
CancelCharge(
    chargeId string,
    request data.CreateCancelChargeRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetChargeResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `chargeId` | `string` | Template, Required | Charge id |
| `request` | [`data.CreateCancelChargeRequest`](../../doc/models/create-cancel-charge-request.md) | Body, Optional | Request for cancelling a charge |
| `idempotencyKey` | `string` | Header, Optional | - |

## Response Type

[`data.GetChargeResponse`](../../doc/models/get-charge-response.md)

## Example Usage

```go
chargeId := "charge_id8"
apiResponse, err := chargesController.CancelCharge(chargeId, request, idempotencyKey)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create Charge

Creates a new charge

```go
CreateCharge(
    request data.CreateChargeRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetChargeResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `request` | [`data.CreateChargeRequest`](../../doc/models/create-charge-request.md) | Body, Required | Request for creating a charge |
| `idempotencyKey` | `string` | Header, Optional | - |

## Response Type

[`data.GetChargeResponse`](../../doc/models/get-charge-response.md)

## Example Usage

```go
requestCustomerAddressMetadata := map[string]string{"Liquid error: Object of type 'System.Collections.Generic.KeyValuePair`2[System.String,System.Object]' cannot be converted to type 'System.String'.", "Liquid error: Object of type 'System.Collections.Generic.KeyValuePair`2[System.String,System.Object]' cannot be converted to type 'System.String'."}
requestCustomerAddress := data.CreateAddressRequest { 
    Street: "street2",
    Number: "number0",
    ZipCode: "zip_code6",
    Neighborhood: "neighborhood8",
    City: "city2",
    State: "state8",
    Country: "country6",
    Complement: "complement8",
    Metadata: requestCustomerAddressMetadata,
    Line1: "line_16",
    Line2: "line_20",
}

requestCustomerMetadata := map[string]string{"Liquid error: Object of type 'System.Collections.Generic.KeyValuePair`2[System.String,System.Object]' cannot be converted to type 'System.String'.", "Liquid error: Object of type 'System.Collections.Generic.KeyValuePair`2[System.String,System.Object]' cannot be converted to type 'System.String'.", "Liquid error: Object of type 'System.Collections.Generic.KeyValuePair`2[System.String,System.Object]' cannot be converted to type 'System.String'."}
requestCustomerPhones := data.CreatePhonesRequest { }

requestCustomer := data.CreateCustomerRequest { 
    Name: "{\n    "name": "Tony Stark"\n}",
    Email: "email0",
    Document: "document0",
    Type: "type4",
    Address: requestCustomerAddress,
    Metadata: requestCustomerMetadata,
    Phones: requestCustomerPhones,
    Code: "code4",
}

requestPayment := data.CreatePaymentRequest { 
    PaymentMethod: "payment_method2",
}

requestMetadata := map[string]string{"Liquid error: Object of type 'System.Collections.Generic.KeyValuePair`2[System.String,System.Object]' cannot be converted to type 'System.String'."}
requestAntifraudClearsale := data.CreateClearSaleRequest { 
    CustomSla: 52,
}

requestAntifraud := data.CreateAntifraudRequest { 
    Type: "type0",
    Clearsale: requestAntifraudClearsale,
}

request := data.CreateChargeRequest { 
    Code: "code4",
    Amount: 242,
    CustomerId: "customer_id4",
    Customer: requestCustomer,
    Payment: requestPayment,
    Metadata: requestMetadata,
    Antifraud: requestAntifraud,
    OrderId: "order_id0",
}

apiResponse, err := chargesController.CreateCharge(request, idempotencyKey)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Confirm Payment

```go
ConfirmPayment(
    chargeId string,
    request data.CreateConfirmPaymentRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetChargeResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `chargeId` | `string` | Template, Required | - |
| `request` | [`data.CreateConfirmPaymentRequest`](../../doc/models/create-confirm-payment-request.md) | Body, Optional | Request for confirm payment |
| `idempotencyKey` | `string` | Header, Optional | - |

## Response Type

[`data.GetChargeResponse`](../../doc/models/get-charge-response.md)

## Example Usage

```go
chargeId := "charge_id8"
apiResponse, err := chargesController.ConfirmPayment(chargeId, request, idempotencyKey)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

