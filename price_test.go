package cryptocomparego

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestPriceList(t *testing.T) {
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

	acct, _, err := client.Price.List(ctx)
	if err != nil {
		t.Errorf("Price.List returned error: %v", err)
	}

	expected := []Price{{"BTC", 0.04502}, {"EUR", 313.82}, {"USD", 368.87}}

	if !reflect.DeepEqual(acct, expected) {
		t.Errorf("Price.List returned %+v, expected %+v", acct, expected)
	}
}
