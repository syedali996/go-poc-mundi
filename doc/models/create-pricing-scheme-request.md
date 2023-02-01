
# Create Pricing Scheme Request

Request for creating a pricing scheme

## Structure

`CreatePricingSchemeRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `schemeType` | `string` | Required | Scheme type |
| `priceBrackets` | [`[]data.CreatePriceBracketRequest`](../../doc/models/create-price-bracket-request.md) | Optional | Price brackets |
| `price` | `int` | Optional | Price |
| `minimumPrice` | `int` | Optional | Minimum price |
| `percentage` | `float64` | Optional | percentual value used in pricing_scheme Percent |

## Example (as JSON)

```json
{
  "scheme_type": "scheme_type0",
  "price_brackets": null,
  "price": null,
  "minimum_price": null,
  "percentage": null
}
```

