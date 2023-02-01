
# Get Subscription Item Response

## Structure

`GetSubscriptionItemResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `*string` | Optional | - |
| `description` | `*string` | Optional | - |
| `status` | `*string` | Optional | - |
| `createdAt` | `*time.Time` | Optional | - |
| `updatedAt` | `*time.Time` | Optional | - |
| `pricingScheme` | [`*data.GetPricingSchemeResponse`](../../doc/models/get-pricing-scheme-response.md) | Optional | - |
| `discounts` | [`*[]data.GetDiscountResponse`](../../doc/models/get-discount-response.md) | Optional | - |
| `increments` | [`*[]data.GetIncrementResponse`](../../doc/models/get-increment-response.md) | Optional | - |
| `subscription` | [`*data.GetSubscriptionResponse`](../../doc/models/get-subscription-response.md) | Optional | - |
| `name` | `*string` | Optional | Item name |
| `quantity` | `*int` | Optional | - |
| `cycles` | `*int` | Optional | - |
| `deletedAt` | `*time.Time` | Optional | - |

## Example (as JSON)

```json
{
  "id": null,
  "description": null,
  "status": null,
  "created_at": null,
  "updated_at": null,
  "pricing_scheme": null,
  "discounts": null,
  "increments": null,
  "subscription": null,
  "name": null,
  "quantity": null,
  "cycles": null,
  "deleted_at": null
}
```

