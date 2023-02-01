# Transfers

```go
transfersController := client.TransfersController
```

## Class Name

`TransfersController`

## Methods

* [Get Transfer by Id](../../doc/controllers/transfers.md#get-transfer-by-id)
* [Create Transfer](../../doc/controllers/transfers.md#create-transfer)
* [Get Transfers](../../doc/controllers/transfers.md#get-transfers)


# Get Transfer by Id

```go
GetTransferById(
    transferId string) (
    https.ApiResponse[data.GetTransfer],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `transferId` | `string` | Template, Required | - |

## Response Type

[`data.GetTransfer`](../../doc/models/get-transfer.md)

## Example Usage

```go
transferId := "transfer_id6"
apiResponse, err := transfersController.GetTransferById(transferId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create Transfer

```go
CreateTransfer(
    request data.CreateTransfer) (
    https.ApiResponse[data.GetTransfer],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `request` | [`data.CreateTransfer`](../../doc/models/create-transfer.md) | Body, Required | - |

## Response Type

[`data.GetTransfer`](../../doc/models/get-transfer.md)

## Example Usage

```go
request := data.CreateTransfer { 
    Amount: 242,
    SourceId: "source_id0",
    TargetId: "target_id6",
}

apiResponse, err := transfersController.CreateTransfer(request)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Transfers

Gets all transfers

```go
GetTransfers() (
    https.ApiResponse[data.ListTransfers],
    error)
```

## Response Type

[`data.ListTransfers`](../../doc/models/list-transfers.md)

## Example Usage

```go
apiResponse, err := transfersController.GetTransfers()
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

