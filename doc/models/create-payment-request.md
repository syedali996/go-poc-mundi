
# Create Payment Request

Payment data

## Structure

`CreatePaymentRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `paymentMethod` | `string` | Required | Payment method |
| `creditCard` | [`data.CreateCreditCardPaymentRequest`](../../doc/models/create-credit-card-payment-request.md) | Optional | Settings for credit card payment |
| `debitCard` | [`data.CreateDebitCardPaymentRequest`](../../doc/models/create-debit-card-payment-request.md) | Optional | Settings for debit card payment |
| `boleto` | [`data.CreateBoletoPaymentRequest`](../../doc/models/create-boleto-payment-request.md) | Optional | Settings for boleto payment |
| `currency` | `string` | Optional | Currency. Must be informed using 3 characters |
| `voucher` | [`data.CreateVoucherPaymentRequest`](../../doc/models/create-voucher-payment-request.md) | Optional | Settings for voucher payment |
| `split` | [`[]data.CreateSplitRequest`](../../doc/models/create-split-request.md) | Optional | Splits |
| `bankTransfer` | [`data.CreateBankTransferPaymentRequest`](../../doc/models/create-bank-transfer-payment-request.md) | Optional | Settings for bank transfer payment |
| `gatewayAffiliationId` | `string` | Optional | Gateway affiliation code |
| `amount` | `int` | Optional | The amount of the payment, in cents |
| `checkout` | [`data.CreateCheckoutPaymentRequest`](../../doc/models/create-checkout-payment-request.md) | Optional | Settings for checkout payment |
| `customerId` | `string` | Optional | Customer Id |
| `customer` | [`data.CreateCustomerRequest`](../../doc/models/create-customer-request.md) | Optional | Customer |
| `metadata` | `map[string]string` | Optional | Metadata |
| `cash` | [`data.CreateCashPaymentRequest`](../../doc/models/create-cash-payment-request.md) | Optional | Settings for cash payment |
| `privateLabel` | [`data.CreatePrivateLabelPaymentRequest`](../../doc/models/create-private-label-payment-request.md) | Optional | Settings for private label payment |
| `pix` | [`data.CreatePixPaymentRequest`](../../doc/models/create-pix-payment-request.md) | Optional | Settings for pix payment |

## Example (as JSON)

```json
{
  "payment_method": "payment_method0",
  "credit_card": null,
  "debit_card": null,
  "boleto": null,
  "currency": null,
  "voucher": null,
  "split": null,
  "bank_transfer": null,
  "gateway_affiliation_id": null,
  "amount": null,
  "checkout": null,
  "customer_id": null,
  "customer": null,
  "metadata": null,
  "cash": null,
  "private_label": null,
  "pix": null
}
```

