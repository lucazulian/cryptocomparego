package cryptocomparego

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestPriceList_NilPriceRequest(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/data/price", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		response := `
		{
			"BTC": 0.04502,
			"EUR": 313.82,
			"USD": 368.87
		}`

		fmt.Fprint(w, response)
	})

	acct, _, err := client.Price.List(ctx, nil)
	if err != nil {
		t.Errorf("Price.List returned error: %v", err)
	}

	expected := []Price{{"BTC", 0.04502}, {"EUR", 313.82}, {"USD", 368.87}}

	if !reflect.DeepEqual(acct, expected) {
		t.Errorf("Price.List returned %+v, expected %+v", acct, expected)
	}
}

func TestPriceList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/data/price", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("fsym") != "ETH" {
			t.Errorf("Price.List did not request with a fsym parameter")
		}

		if r.URL.Query().Get("tsyms") != "BTC,USD,EUR" {
			t.Errorf("Price.List did not request with a tsyms parameter")
		}

		testMethod(t, r, http.MethodGet)

		response := `
		{
			"BTC": 0.04502,
			"EUR": 313.82,
			"USD": 368.87
		}`

		fmt.Fprint(w, response)
	})

	priceRequest := &PriceRequest{Fsym: "ETH", Tsyms: []string{"BTC", "USD", "EUR"}}

	acct, _, err := client.Price.List(ctx, priceRequest)
	if err != nil {
		t.Errorf("Price.List returned error: %v", err)
	}

	expected := []Price{{"BTC", 0.04502}, {"EUR", 313.82}, {"USD", 368.87}}

	if !reflect.DeepEqual(acct, expected) {
		t.Errorf("Price.List returned %+v, expected %+v", acct, expected)
	}
}
