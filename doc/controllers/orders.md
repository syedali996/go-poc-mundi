# Orders

```go
ordersController := client.OrdersController
```

## Class Name

`OrdersController`

## Methods

* [Get Orders](../../doc/controllers/orders.md#get-orders)
* [Update Order Item](../../doc/controllers/orders.md#update-order-item)
* [Delete All Order Items](../../doc/controllers/orders.md#delete-all-order-items)
* [Delete Order Item](../../doc/controllers/orders.md#delete-order-item)
* [Close Order](../../doc/controllers/orders.md#close-order)
* [Create Order](../../doc/controllers/orders.md#create-order)
* [Create Order Item](../../doc/controllers/orders.md#create-order-item)
* [Get Order Item](../../doc/controllers/orders.md#get-order-item)
* [Update Order Metadata](../../doc/controllers/orders.md#update-order-metadata)
* [Get Order](../../doc/controllers/orders.md#get-order)


# Get Orders

Gets all orders

```go
GetOrders(
    page int,
    size int,
    code string,
    status string,
    createdSince time.Time,
    createdUntil time.Time,
    customerId string) (
    https.ApiResponse[data.ListOrderResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `page` | `int` | Query, Optional | Page number |
| `size` | `int` | Query, Optional | Page size |
| `code` | `string` | Query, Optional | Filter for order's code |
| `status` | `string` | Query, Optional | Filter for order's status |
| `createdSince` | `time.Time` | Query, Optional | Filter for order's creation date start range |
| `createdUntil` | `time.Time` | Query, Optional | Filter for order's creation date end range |
| `customerId` | `string` | Query, Optional | Filter for order's customer id |

## Response Type

[`data.ListOrderResponse`](../../doc/models/list-order-response.md)

## Example Usage

```go
apiResponse, err := ordersController.GetOrders(page, size, code, status, createdSince, createdUntil, customerId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Order Item

```go
UpdateOrderItem(
    orderId string,
    itemId string,
    request data.UpdateOrderItemRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetOrderItemResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orderId` | `string` | Template, Required | Order Id |
| `itemId` | `string` | Template, Required | Item Id |
| `request` | [`data.UpdateOrderItemRequest`](../../doc/models/update-order-item-request.md) | Body, Required | Item Model |
| `idempotencyKey` | `string` | Header, Optional | - |

## Response Type

[`data.GetOrderItemResponse`](../../doc/models/get-order-item-response.md)

## Example Usage

```go
orderId := "orderId2"
itemId := "itemId8"
request := data.UpdateOrderItemRequest { 
    Amount: 242,
    Description: "description6",
    Quantity: 100,
    Category: "category4",
}

apiResponse, err := ordersController.UpdateOrderItem(orderId, itemId, request, idempotencyKey)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Delete All Order Items

```go
DeleteAllOrderItems(
    orderId string,
    idempotencyKey string) (
    https.ApiResponse[data.GetOrderResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orderId` | `string` | Template, Required | Order Id |
| `idempotencyKey` | `string` | Header, Optional | - |

## Response Type

[`data.GetOrderResponse`](../../doc/models/get-order-response.md)

## Example Usage

```go
orderId := "orderId2"
apiResponse, err := ordersController.DeleteAllOrderItems(orderId, idempotencyKey)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Delete Order Item

```go
DeleteOrderItem(
    orderId string,
    itemId string,
    idempotencyKey string) (
    https.ApiResponse[data.GetOrderItemResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orderId` | `string` | Template, Required | Order Id |
| `itemId` | `string` | Template, Required | Item Id |
| `idempotencyKey` | `string` | Header, Optional | - |

## Response Type

[`data.GetOrderItemResponse`](../../doc/models/get-order-item-response.md)

## Example Usage

```go
orderId := "orderId2"
itemId := "itemId8"
apiResponse, err := ordersController.DeleteOrderItem(orderId, itemId, idempotencyKey)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Close Order

```go
CloseOrder(
    id string,
    request data.UpdateOrderStatusRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetOrderResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `string` | Template, Required | Order Id |
| `request` | [`data.UpdateOrderStatusRequest`](../../doc/models/update-order-status-request.md) | Body, Required | Update Order Model |
| `idempotencyKey` | `string` | Header, Optional | - |

## Response Type

[`data.GetOrderResponse`](../../doc/models/get-order-response.md)

## Example Usage

```go
id := "id0"
request := data.UpdateOrderStatusRequest { 
    Status: "status8",
}

apiResponse, err := ordersController.CloseOrder(id, request, idempotencyKey)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create Order

Creates a new Order

```go
CreateOrder(
    body data.CreateOrderRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetOrderResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`data.CreateOrderRequest`](../../doc/models/create-order-request.md) | Body, Required | Request for creating an order |
| `idempotencyKey` | `string` | Header, Optional | - |

## Response Type

[`data.GetOrderResponse`](../../doc/models/get-order-response.md)

## Example Usage

```go
bodyItems := []data.CreateOrderItemRequest {}

bodyItems0 := data.CreateOrderItemRequest { 
    Amount: 101,
    Description: "description3",
    Quantity: 215,
    Category: "category1",
}

bodyItems = append(bodyItems, bodyItems0)

bodyItems1 := data.CreateOrderItemRequest { 
    Amount: 102,
    Description: "description4",
    Quantity: 216,
    Category: "category2",
}

bodyItems = append(bodyItems, bodyItems1)

bodyItems2 := data.CreateOrderItemRequest { 
    Amount: 103,
    Description: "description5",
    Quantity: 217,
    Category: "category3",
}

bodyItems = append(bodyItems, bodyItems2)

bodyCustomerAddressMetadata := map[string]string{"Liquid error: Object of type 'System.Collections.Generic.KeyValuePair`2[System.String,System.Object]' cannot be converted to type 'System.String'.", "Liquid error: Object of type 'System.Collections.Generic.KeyValuePair`2[System.String,System.Object]' cannot be converted to type 'System.String'."}
bodyCustomerAddress := data.CreateAddressRequest { 
    Street: "street0",
    Number: "number8",
    ZipCode: "zip_code4",
    Neighborhood: "neighborhood6",
    City: "city0",
    State: "state6",
    Country: "country4",
    Complement: "complement6",
    Metadata: bodyCustomerAddressMetadata,
    Line1: "line_16",
    Line2: "line_28",
}

bodyCustomerMetadata := map[string]string{"Liquid error: Object of type 'System.Collections.Generic.KeyValuePair`2[System.String,System.Object]' cannot be converted to type 'System.String'.", "Liquid error: Object of type 'System.Collections.Generic.KeyValuePair`2[System.String,System.Object]' cannot be converted to type 'System.String'."}
bodyCustomerPhones := data.CreatePhonesRequest { }

bodyCustomer := data.CreateCustomerRequest { 
    Name: "{\n    "name": "Tony Stark"\n}",
    Email: "email2",
    Document: "document2",
    Type: "type6",
    Address: bodyCustomerAddress,
    Metadata: bodyCustomerMetadata,
    Phones: bodyCustomerPhones,
    Code: "code2",
}

bodyPayments := []data.CreatePaymentRequest {}

bodyPayments0 := data.CreatePaymentRequest { 
    PaymentMethod: "payment_method0",
}

bodyPayments = append(bodyPayments, bodyPayments0)

bodyPayments1 := data.CreatePaymentRequest { 
    PaymentMethod: "payment_method9",
}

bodyPayments = append(bodyPayments, bodyPayments1)

bodyMetadata := map[string]string{"Liquid error: Object of type 'System.Collections.Generic.KeyValuePair`2[System.String,System.Object]' cannot be converted to type 'System.String'.", "Liquid error: Object of type 'System.Collections.Generic.KeyValuePair`2[System.String,System.Object]' cannot be converted to type 'System.String'."}
body := data.CreateOrderRequest { 
    Items: bodyItems,
    Customer: bodyCustomer,
    Payments: bodyPayments,
    Code: "code4",
    CustomerId: "customer_id4",
    Metadata: bodyMetadata,
    Closed: true,
}

apiResponse, err := ordersController.CreateOrder(body, idempotencyKey)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create Order Item

```go
CreateOrderItem(
    orderId string,
    request data.CreateOrderItemRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetOrderItemResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orderId` | `string` | Template, Required | Order Id |
| `request` | [`data.CreateOrderItemRequest`](../../doc/models/create-order-item-request.md) | Body, Required | Order Item Model |
| `idempotencyKey` | `string` | Header, Optional | - |

## Response Type

[`data.GetOrderItemResponse`](../../doc/models/get-order-item-response.md)

## Example Usage

```go
orderId := "orderId2"
request := data.CreateOrderItemRequest { 
    Amount: 242,
    Description: "description6",
    Quantity: 100,
    Category: "category4",
}

apiResponse, err := ordersController.CreateOrderItem(orderId, request, idempotencyKey)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Order Item

```go
GetOrderItem(
    orderId string,
    itemId string) (
    https.ApiResponse[data.GetOrderItemResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orderId` | `string` | Template, Required | Order Id |
| `itemId` | `string` | Template, Required | Item Id |

## Response Type

[`data.GetOrderItemResponse`](../../doc/models/get-order-item-response.md)

## Example Usage

```go
orderId := "orderId2"
itemId := "itemId8"
apiResponse, err := ordersController.GetOrderItem(orderId, itemId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Order Metadata

Updates the metadata from an order

```go
UpdateOrderMetadata(
    orderId string,
    request data.UpdateMetadataRequest,
    idempotencyKey string) (
    https.ApiResponse[data.GetOrderResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orderId` | `string` | Template, Required | The order id |
| `request` | [`data.UpdateMetadataRequest`](../../doc/models/update-metadata-request.md) | Body, Required | Request for updating the order metadata |
| `idempotencyKey` | `string` | Header, Optional | - |

## Response Type

[`data.GetOrderResponse`](../../doc/models/get-order-response.md)

## Example Usage

```go
orderId := "order_id6"
requestMetadata := map[string]string{"Liquid error: Object of type 'System.Collections.Generic.KeyValuePair`2[System.String,System.Object]' cannot be converted to type 'System.String'."}
request := data.UpdateMetadataRequest { 
    Metadata: requestMetadata,
}

apiResponse, err := ordersController.UpdateOrderMetadata(orderId, request, idempotencyKey)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Order

Gets an order

```go
GetOrder(
    orderId string) (
    https.ApiResponse[data.GetOrderResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orderId` | `string` | Template, Required | Order id |

## Response Type

[`data.GetOrderResponse`](../../doc/models/get-order-response.md)

## Example Usage

```go
orderId := "order_id6"
apiResponse, err := ordersController.GetOrder(orderId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

