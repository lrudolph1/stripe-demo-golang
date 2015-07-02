package main

import (
	"fmt"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"github.com/stripe/stripe-go/currency"
	"log"
	"net/http"
)

const (
	AmountToCharge uint64 = 100
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func createDebit(token string, amount uint64, description string) *stripe.Charge {
	stripe.Key = "sk_test_jhNUaFdOfykMNlXWZMmtUdwy" // get your secret test key from 
	// https://dashboard.stripe.com/account/apikeys and place it here 

	params := &stripe.ChargeParams{
		Amount:   amount,
		Currency: currency.USD,
		Desc: "test charge",
	}

	// obtain token from stripe
	params.SetSource(token)

	ch, err := charge.New(params)

	if err != nil {
		log.Fatalf("error while trying to charge a cc", err)
	}

	log.Printf("debit created successfully %v\n", ch.ID)

	return ch
}

func debitsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		createDebit(r.FormValue("stripeToken"), AmountToCharge, "testing charge description!")
		fmt.Fprint(w, "successful payment.")
	}
}


func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/debits", debitsHandler)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8081", nil)
}
