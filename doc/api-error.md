
# ApiError

Thrown when the HTTP status code is not okay.

The ApiError implements error interface.

## Properties

| Name | Type | Description |
|  --- | --- | --- |
| Request | http.Request | Original request that resulted in this response. |
| StatusCode | int | Response status code. |
| Headers | map[string]string | Response headers. |
| Body | string | Original body from the response. |

