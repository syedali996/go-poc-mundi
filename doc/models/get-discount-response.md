
# Get Discount Response

Response object for getting a discount

## Structure

`GetDiscountResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `*string` | Optional | - |
| `value` | `*float64` | Optional | - |
| `discountType` | `*string` | Optional | - |
| `status` | `*string` | Optional | - |
| `createdAt` | `*time.Time` | Optional | - |
| `cycles` | `*int` | Optional | - |
| `deletedAt` | `*time.Time` | Optional | - |
| `description` | `*string` | Optional | - |
| `subscription` | [`*data.GetSubscriptionResponse`](../../doc/models/get-subscription-response.md) | Optional | - |
| `subscriptionItem` | [`*data.GetSubscriptionItemResponse`](../../doc/models/get-subscription-item-response.md) | Optional | The subscription item |

## Example (as JSON)

```json
{
  "id": null,
  "value": null,
  "discount_type": null,
  "status": null,
  "created_at": null,
  "cycles": null,
  "deleted_at": null,
  "description": null,
  "subscription": null,
  "subscription_item": null
}
```

