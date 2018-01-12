package cryptocomparego

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestFormattedQueryStringNilPriceMultiRequest(t *testing.T) {

	priceMultiRequest := NewPriceMultiRequest([]string{}, nil)
	acct := priceMultiRequest.FormattedQueryString("/data/pricemulti")

	expected := "/data/pricemulti?e=CCCAGG&sign=false&tryConversion=true"

	if acct != expected {
		t.Errorf("PriceMultiRequest.FormattedQueryString returned %+v, expected %+v", acct, expected)
	}
}

func TestFormattedQueryStringPriceMultiRequest(t *testing.T) {

	priceMultiRequest := NewPriceMultiRequest([]string{"BTC", "ETH"}, []string{"BTC", "USD", "EUR"})
	acct := priceMultiRequest.FormattedQueryString("/data/pricemulti")

	expected := "/data/pricemulti?e=CCCAGG&fsyms=BTC%2CETH&sign=false&tryConversion=true&tsyms=BTC%2CUSD%2CEUR"

	if acct != expected {
		t.Errorf("PricMultieRequest.FormattedQueryString returned %+v, expected %+v", acct, expected)
	}
}

func TestPriceMultiList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/data/pricemulti", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		if r.URL.Query().Get("fsyms") != "BTC,ETH" || r.URL.Query().Get("tsyms") != "BTC,USD,EUR" {
			t.Errorf("Price.List did not request the correct fsym or tsyms")
		}

		response := `
		{
			"BTC": {
				"BTC": 1,
				"USD": 14842.32,
				"EUR": 13476.09
			},
			"ETH": {
				"BTC": 0.02723,
				"USD": 409.12,
				"EUR": 367
			}
		}`

		fmt.Fprint(w, response)
	})

	priceMultiRequest := &PriceMultiRequest{Fsyms: []string{"BTC", "ETH"}, Tsyms: []string{"BTC", "USD", "EUR"}}

	acct, _, err := client.PriceMulti.List(ctx, priceMultiRequest)
	if err != nil {
		t.Errorf("PriceMulti.List returned error: %v", err)
	}

	expected := []PriceMulti{
		{"BTC", []Price{{"BTC", 1}, {"EUR", 13476.09}, {"USD", 14842.32}}},
		{"ETH", []Price{{"BTC", 0.02723}, {"EUR", 367}, {"USD", 409.12}}},
	}

	if !reflect.DeepEqual(acct, expected) {
		t.Errorf("PriceMulti.List returned %+v, expected %+v", acct, expected)
	}
}
