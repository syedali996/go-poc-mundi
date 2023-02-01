
# Create Subscription Split Request

## Structure

`CreateSubscriptionSplitRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `enabled` | `bool` | Required | Defines if the split is enabled |
| `rules` | [`[]data.CreateSplitRequest`](../../doc/models/create-split-request.md) | Required | Split |

## Example (as JSON)

```json
{
  "enabled": false,
  "rules": [
    {
      "type": "type6",
      "amount": 210,
      "recipient_id": "recipient_id6",
      "options": null,
      "split_rule_id": null
    }
  ]
}
```

