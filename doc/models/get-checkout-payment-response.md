
# Get Checkout Payment Response

Resposta das configurações de pagamento do checkout

## Structure

`GetCheckoutPaymentResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `*string` | Optional | - |
| `amount` | `*int` | Optional | Valor em centavos |
| `defaultPaymentMethod` | `*string` | Optional | Meio de pagamento padrão no checkout |
| `successUrl` | `*string` | Optional | Url de redirecionamento de sucesso após o checkou |
| `paymentUrl` | `*string` | Optional | Url para pagamento usando o checkout |
| `gatewayAffiliationId` | `*string` | Optional | Código da afiliação onde o pagamento será processado no gateway |
| `acceptedPaymentMethods` | `*[]string` | Optional | Meios de pagamento aceitos no checkout |
| `status` | `*string` | Optional | Status do checkout |
| `skipCheckoutSuccessPage` | `*bool` | Optional | Pular tela de sucesso pós-pagamento? |
| `createdAt` | `*time.Time` | Optional | Data de criação |
| `updatedAt` | `*time.Time` | Optional | Data de atualização |
| `canceledAt` | `*time.Time` | Optional | Data de cancelamento |
| `customerEditable` | `*bool` | Optional | Torna o objeto customer editável |
| `customer` | [`*data.GetCustomerResponse`](../../doc/models/get-customer-response.md) | Optional | Dados do comprador |
| `billingaddress` | [`*data.GetAddressResponse`](../../doc/models/get-address-response.md) | Optional | Dados do endereço de cobrança |
| `creditCard` | [`*data.GetCheckoutCreditCardPaymentResponse`](../../doc/models/get-checkout-credit-card-payment-response.md) | Optional | Configurações de cartão de crédito |
| `boleto` | [`*data.GetCheckoutBoletoPaymentResponse`](../../doc/models/get-checkout-boleto-payment-response.md) | Optional | Configurações de boleto |
| `billingAddressEditable` | `*bool` | Optional | Indica se o billing address poderá ser editado |
| `shipping` | [`*data.GetShippingResponse`](../../doc/models/get-shipping-response.md) | Optional | Configurações  de entrega |
| `shippable` | `*bool` | Optional | Indica se possui entrega |
| `closedAt` | `*time.Time` | Optional | Data de fechamento |
| `expiresAt` | `*time.Time` | Optional | Data de expiração |
| `currency` | `*string` | Optional | Moeda |
| `debitCard` | [`*data.GetCheckoutDebitCardPaymentResponse`](../../doc/models/get-checkout-debit-card-payment-response.md) | Optional | Configurações de cartão de débito |
| `bankTransfer` | [`*data.GetCheckoutBankTransferPaymentResponse`](../../doc/models/get-checkout-bank-transfer-payment-response.md) | Optional | Bank transfer payment response |
| `acceptedBrands` | `*[]string` | Optional | Accepted Brands |
| `pix` | [`*data.GetCheckoutPixPaymentResponse`](../../doc/models/get-checkout-pix-payment-response.md) | Optional | Pix payment response |

## Example (as JSON)

```json
{
  "id": null,
  "amount": null,
  "default_payment_method": null,
  "success_url": null,
  "payment_url": null,
  "gateway_affiliation_id": null,
  "accepted_payment_methods": null,
  "status": null,
  "skip_checkout_success_page": null,
  "created_at": null,
  "updated_at": null,
  "canceled_at": null,
  "customer_editable": null,
  "customer": null,
  "billingaddress": null,
  "credit_card": null,
  "boleto": null,
  "billing_address_editable": null,
  "shipping": null,
  "shippable": null,
  "closed_at": null,
  "expires_at": null,
  "currency": null,
  "debit_card": null,
  "bank_transfer": null,
  "accepted_brands": null,
  "pix": null
}
```

