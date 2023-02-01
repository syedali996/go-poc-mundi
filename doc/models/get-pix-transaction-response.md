
# Get Pix Transaction Response

Response object when getting a pix transaction

## Structure

`GetPixTransactionResponse`

## Inherits From

[`GetTransactionResponse`](../../doc/models/get-transaction-response.md)

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `qrCode` | `*string` | Optional | - |
| `qrCodeUrl` | `*string` | Optional | - |
| `expiresAt` | `*time.Time` | Optional | - |
| `additionalInformation` | [`*[]data.PixAdditionalInformation`](../../doc/models/pix-additional-information.md) | Optional | - |
| `endToEndId` | `*string` | Optional | - |
| `payer` | [`*data.GetPixPayerResponse`](../../doc/models/get-pix-payer-response.md) | Optional | - |

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
  "transaction_type": "pix",
  "id": null,
  "gateway_response": null,
  "antifraud_response": null,
  "metadata": null,
  "split": null,
  "interest": null,
  "fine": null,
  "max_days_to_pay_past_due": null,
  "qr_code": null,
  "qr_code_url": null,
  "expires_at": null,
  "additional_information": null,
  "end_to_end_id": null,
  "payer": null
}
```

