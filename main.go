package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

type PayRequest struct {
	OrderID string `json:"order_id"`
}

func configure() {
	stripe.Key = ""
	stripe.SetBackend(stripe.APIBackend,
		stripe.GetBackendWithConfig(stripe.APIBackend,
			&stripe.BackendConfig{
				LogLevel:          5,
				MaxNetworkRetries: 3,
				HTTPClient:        &http.Client{Timeout: 10 * time.Second},
			},
		),
	)
}

func pay(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var req PayRequest
	err := decoder.Decode(&req)
	if err != nil {
		panic(err)
	}

	chargeParams := &stripe.ChargeParams{
		Amount:   stripe.Int64(int64(1000)),
		Currency: stripe.String(string(stripe.CurrencyCHF)),
		Customer: stripe.String("cus_CGdSFwA12HrcS8"),
		Destination: &stripe.DestinationParams{
			Amount:  stripe.Int64(int64(1000)),
			Account: stripe.String("acct_1CroMIC3GdU9y0PX"),
		},
	}

	chargeParams.SetSource("src_1CuhE6C1oEzKtlmej7JBjlwg")
	chargeParams.IdempotencyKey = stripe.String(req.OrderID)

	charge, err := charge.New(chargeParams)
	if err != nil {
		fmt.Println("Stripe charge failed", err)
		return
	}
	fmt.Println("Stripe charge OK", charge.ID)
}

func main() {
	configure()
	http.HandleFunc("/charges/", pay)
	if err := http.ListenAndServe("localhost:8090", nil); err != nil {
		panic(err)
	}
}
