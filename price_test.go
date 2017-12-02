package cryptocomparego

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)


func TestFormattedQueryString_NilPriceRequest(t *testing.T) {

	priceRequest := NewPriceRequest("", nil)
	acct := priceRequest.FormattedQueryString("/data/price")

	expected := "/data/price?e=CCCAGG&sign=false&tryConversion=true"

	if acct != expected {
		t.Errorf("PriceRequest.FormattedQueryString returned %+v, expected %+v", acct, expected)
	}
}

func TestFormattedQueryString(t *testing.T) {

	priceRequest := NewPriceRequest("ETH", []string{"BTC", "USD", "EUR"})
	acct := priceRequest.FormattedQueryString("/data/price")

	expected := "/data/price?fsym=ETH&tsyms=BTC,USD,EUR&e=CCCAGG&sign=false&tryConversion=true"

	if acct != expected {
		t.Errorf("PriceRequest.FormattedQueryString returned %+v, expected %+v", acct, expected)
	}
}

func TestPriceList_NilPriceRequest(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/data/price", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		if r.URL.Query().Get("fsym") != "" || r.URL.Query().Get("tsyms") != "" {
			t.Errorf("Price.List did not request the correct fsym or tsyms")
		}

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

func TestPriceList_WrongPriceRequest(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/data/price", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		if r.URL.Query().Get("fsym") != "ETH" || r.URL.Query().Get("tsyms") != "" {
			t.Errorf("Price.List did not request the correct fsym or tsyms")
		}

		response := `
		{
			"Response": "Error",
			"Message": "tsyms param seems to be missing.",
			"Type": 1,
			"Aggregated": false,
			"Data": []
		}`

		fmt.Fprint(w, response)
	})

	priceRequest := NewPriceRequest("ETH", nil)

	acct, _, err := client.Price.List(ctx, priceRequest)
	if acct != nil {
		t.Errorf("Price.List returned a value: %v", err)
	}

	expected := errors.New("tsyms param seems to be missing.")

	if !reflect.DeepEqual(err, expected) {
		t.Errorf("Price.List returned %+v, expected %+v", acct, expected)
	}
}

func TestPriceList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/data/price", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		if r.URL.Query().Get("fsym") != "ETH" || r.URL.Query().Get("tsyms") != "BTC,USD,EUR" {
			t.Errorf("Price.List did not request the correct fsym or tsyms")
		}

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
