# stripe-go-invalid-char

1. Edit `main.go` with your stripe private key, a source, a destination account and a customer.
2. `go run main.go`
3. `yarn && node index.js`

The script will trigger two charge requests with the same idempotency key.

Expected output:
```
2018/08/04 00:08:16 Requesting POST api.stripe.com/v1/charges
2018/08/04 00:08:17 Requesting POST api.stripe.com/v1/charges
2018/08/04 00:08:18 Request completed in 1.255210765s (retry: 0)
2018/08/04 00:08:18 Request failed with: {"error":{"code":"idempotency_key_in_use","doc_url":"https://stripe.com/docs/error-codes/idempotency-key-in-use","message":"There is currently another in-progress request using this Idempotent Key (that probably means you submitted twice, and the other request is still going through): 770434815658t5litna. Please try again later.","type":"invalid_request_error"}} (error: <nil>)
2018/08/04 00:08:18 Initiating retry 1 for request POST api.stripe.com/v1/charges after sleeping 500ms
2018/08/04 00:08:19 Request completed in 242.108105ms (retry: 1)
Stripe charge failed invalid character '<' looking for beginning of value
2018/08/04 00:08:19 Request completed in 3.031320486s (retry: 0)
2018/08/04 00:08:19 Response: {
  "id": "ch_1CvBRVC1oEzKtlmeiOn2lvlv",
 ``` 
 
 Additional information:
 
 One of the two request is expected to end with a 409 and be retried. The retry will result in a 400 (??) status code causing `shouldRetry` to return `false`.
 
 Since `err` is `nil`, the condition `res.StatusCode >= 400` will evaluate to `true`. Finally, the `ResponseToError` function will return an error after trying to `Unmarshal`
 
