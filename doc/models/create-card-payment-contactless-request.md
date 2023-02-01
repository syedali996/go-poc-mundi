
# Create Card Payment Contactless Request

The card payment contactless request

## Structure

`CreateCardPaymentContactlessRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mType` | `string` | Required | The authentication type |
| `applePay` | [`data.CreateApplePayRequest`](../../doc/models/create-apple-pay-request.md) | Optional | The ApplePay encrypted request |
| `googlePay` | [`data.CreateGooglePayRequest`](../../doc/models/create-google-pay-request.md) | Optional | The GooglePay encrypted request |
| `emv` | [`data.CreateEmvDecryptRequest`](../../doc/models/create-emv-decrypt-request.md) | Optional | The Emv encrypted request |

## Example (as JSON)

```json
{
  "type": "type0",
  "apple_pay": null,
  "google_pay": null,
  "emv": null
}
```

