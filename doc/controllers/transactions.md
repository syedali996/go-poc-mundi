# Transactions

```go
transactionsController := client.TransactionsController
```

## Class Name

`TransactionsController`


# Get Transaction

```go
GetTransaction(
    transactionId string) (
    https.ApiResponse[data.GetTransactionResponseInterface],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `transactionId` | `string` | Template, Required | - |

## Response Type

[`data.GetTransactionResponseInterface`](../../doc/models/get-transaction-response.md)

## Example Usage

```go
transactionId := "transaction_id8"
apiResponse, err := transactionsController.GetTransaction(transactionId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

