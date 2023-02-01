
# Create Antifraud Request

## Structure

`CreateAntifraudRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mType` | `string` | Required | - |
| `clearsale` | [`data.CreateClearSaleRequest`](../../doc/models/create-clear-sale-request.md) | Required | - |

## Example (as JSON)

```json
{
  "type": "type0",
  "clearsale": {
    "custom_sla": 178
  }
}
```

