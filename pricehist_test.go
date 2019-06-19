package cryptocomparego

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/lucazulian/cryptocomparego/context"
)

func TestFormattedQueryStringNilPriceHistRequest(t *testing.T) {

	priceHistRequest := NewPriceHistRequest("", nil, 1560918228)
	acct := priceHistRequest.FormattedQueryString("/data/pricehistorical")

	expected := "/data/pricehistorical?calculationType=Close&e=CCCAGG&sign=false&tryConversion=true&ts=1560918228"

	if acct != expected {
		t.Errorf("PriceHistRequest.FormattedQueryString returned %+v, expected %+v", acct, expected)
	}
}

func TestFormattedQueryStringPriceHistRequest(t *testing.T) {

	priceHistRequest := NewPriceHistRequest("ETH", []string{"BTC", "USD", "EUR"}, 1560918228)
	acct := priceHistRequest.FormattedQueryString("/data/pricehistorical")

	expected := "/data/pricehistorical?calculationType=Close&e=CCCAGG&fsym=ETH&sign=false&tryConversion=true&ts=1560918228&tsyms=BTC%2CUSD%2CEUR"

	if acct != expected {
		t.Errorf("PriceRequest.FormattedQueryString returned %+v, expected %+v", acct, expected)
	}
}

func TestPriceListNilPriceHistRequest(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/data/pricehistorical", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		if r.URL.Query().Get("fsym") != "" || r.URL.Query().Get("tsyms") != "" {
			t.Errorf("PriceHist.List did not request the correct fsym or tsyms")
		}

		response := `
		{
			"BTC": 0.04502,
			"EUR": 313.82,
			"USD": 368.87
		}`

		fmt.Fprint(w, response)
	})

	acct, _, err := client.PriceHist.List(ctx, nil)
	if err != nil {
		t.Errorf("PriceHist.List returned error: %v", err)
	}

	expected := []PriceHist{{"BTC", 0.04502}, {"EUR", 313.82}, {"USD", 368.87}}

	if !reflect.DeepEqual(acct, expected) {
		t.Errorf("PriceHist.List returned %+v, expected %+v", acct, expected)
	}
}

func TestPriceListWrongPriceHistRequest(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/data/pricehistorical", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		if r.URL.Query().Get("fsym") != "ETH" || r.URL.Query().Get("tsyms") != "" {
			t.Errorf("PriceHist.List did not request the correct fsym or tsyms")
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

	priceHistRequest := NewPriceHistRequest("ETH", nil, 1560918228)

	acct, _, err := client.PriceHist.List(ctx, priceHistRequest)
	if acct != nil {
		t.Errorf("PriceHist.List returned a value: %v", err)
	}

	expected := errors.New("tsyms param seems to be missing.")

	if !reflect.DeepEqual(err, expected) {
		t.Errorf("PriceHist.List returned %+v, expected %+v", acct, expected)
	}
}

func TestPriceHistList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/data/pricehistorical", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		if r.URL.Query().Get("fsym") != "ETH" || r.URL.Query().Get("tsyms") != "BTC,USD,EUR" {
			t.Errorf("PriceHist.List did not request the correct fsym or tsyms")
		}

		response := `
		{
			"BTC": 0.04502,
			"EUR": 313.82,
			"USD": 368.87
		}`

		fmt.Fprint(w, response)
	})

	priceHistRequest := &PriceHistRequest{Fsym: "ETH", Tsyms: []string{"BTC", "USD", "EUR"}}

	acct, _, err := client.PriceHist.List(ctx, priceHistRequest)
	if err != nil {
		t.Errorf("PriceHist.List returned error: %v", err)
	}

	expected := []PriceHist{{"BTC", 0.04502}, {"EUR", 313.82}, {"USD", 368.87}}

	if !reflect.DeepEqual(acct, expected) {
		t.Errorf("PriceHist.List returned %+v, expected %+v", acct, expected)
	}
}

func TestGetPriceHist(t *testing.T) {
	client := NewClient(nil)
	ctx := context.TODO()
	priceHistRequest := NewPriceHistRequest("ETH", []string{"BTC", "USD", "EUR"}, 1560918228)
	acct, _, err := client.PriceHist.List(ctx, priceHistRequest)

	if err != nil {
		t.Errorf("PriceHist.List returned error: %v", err)
	}
	expected := []PriceHist{{"BTC", 0.029150}, {"EUR", 235.630000}, {"USD", 263.870000}}

	if !reflect.DeepEqual(err, expected) {
		t.Errorf("PriceHist.List returned %+v, expected %+v", acct, expected)
	}
}
