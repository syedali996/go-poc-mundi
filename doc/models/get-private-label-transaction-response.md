
# Get Private Label Transaction Response

Response object for getting a private label transaction

## Structure

`GetPrivateLabelTransactionResponse`

## Inherits From

[`GetTransactionResponse`](../../doc/models/get-transaction-response.md)

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `statementDescriptor` | `*string` | Optional | Text that will appear on the credit card's statement |
| `acquirerName` | `*string` | Optional | Acquirer name |
| `acquirerAffiliationCode` | `*string` | Optional | Aquirer affiliation code |
| `acquirerTid` | `*string` | Optional | Acquirer TID |
| `acquirerNsu` | `*string` | Optional | Acquirer NSU |
| `acquirerAuthCode` | `*string` | Optional | Acquirer authorization code |
| `operationType` | `*string` | Optional | Operation type |
| `card` | [`*data.GetCardResponse`](../../doc/models/get-card-response.md) | Optional | Card data |
| `acquirerMessage` | `*string` | Optional | Acquirer message |
| `acquirerReturnCode` | `*string` | Optional | Acquirer Return Code |
| `installments` | `*int` | Optional | Number of installments |

## Example (as JSON)

```json
{
  "gateway_id": null,
  "amount": null,
  "status": null,
  "success": null,
  "created_at": null,
  "updated_at": null,
  "attempt_count": null,
  "max_attempts": null,
  "splits": null,
  "next_attempt": null,
  "transaction_type": "private_label",
  "id": null,
  "gateway_response": null,
  "antifraud_response": null,
  "metadata": null,
  "split": null,
  "interest": null,
  "fine": null,
  "max_days_to_pay_past_due": null,
  "statement_descriptor": null,
  "acquirer_name": null,
  "acquirer_affiliation_code": null,
  "acquirer_tid": null,
  "acquirer_nsu": null,
  "acquirer_auth_code": null,
  "operation_type": null,
  "card": null,
  "acquirer_message": null,
  "acquirer_return_code": null,
  "installments": null
}
```

