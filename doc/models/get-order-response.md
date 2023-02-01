
# Get Order Response

Response object for getting an Order

## Structure

`GetOrderResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `*string` | Optional | - |
| `code` | `*string` | Optional | - |
| `currency` | `*string` | Optional | - |
| `items` | [`*[]data.GetOrderItemResponse`](../../doc/models/get-order-item-response.md) | Optional | - |
| `customer` | [`*data.GetCustomerResponse`](../../doc/models/get-customer-response.md) | Optional | - |
| `status` | `*string` | Optional | - |
| `createdAt` | `*time.Time` | Optional | - |
| `updatedAt` | `*time.Time` | Optional | - |
| `charges` | [`*[]data.GetChargeResponse`](../../doc/models/get-charge-response.md) | Optional | - |
| `invoiceUrl` | `*string` | Optional | - |
| `shipping` | [`*data.GetShippingResponse`](../../doc/models/get-shipping-response.md) | Optional | - |
| `metadata` | `map[string]*string` | Optional | - |
| `checkouts` | [`*[]data.GetCheckoutPaymentResponse`](../../doc/models/get-checkout-payment-response.md) | Optional | Checkout Payment Settings Response |
| `ip` | `*string` | Optional | Ip address |
| `sessionId` | `*string` | Optional | Session id |
| `location` | [`*data.GetLocationResponse`](../../doc/models/get-location-response.md) | Optional | Location |
| `device` | [`*data.GetDeviceResponse`](../../doc/models/get-device-response.md) | Optional | Device's informations |
| `closed` | `*bool` | Optional | Indicates whether the order is closed |

## Example (as JSON)

```json
{
  "id": null,
  "code": null,
  "currency": null,
  "items": null,
  "customer": null,
  "status": null,
  "created_at": null,
  "updated_at": null,
  "charges": null,
  "invoice_url": null,
  "shipping": null,
  "metadata": null,
  "checkouts": null,
  "ip": null,
  "session_id": null,
  "location": null,
  "device": null,
  "closed": null
}
```

