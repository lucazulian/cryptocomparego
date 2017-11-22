package cryptocomparego

import (
	"testing"
	"fmt"
	"net/http"
	"reflect"
)

func TestCoinListList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/data/coinlist", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		response := `
		{
			"Response": "Success",
			"Message": "Coin list succesfully returned!",
			"BaseImageUrl": "https://www.cryptocompare.com",
			"BaseLinkUrl": "https://www.cryptocompare.com",
			"Data": {
				"LTC": {
					"Id": "3808",
					"Url": "/coins/ltc/overview",
					"ImageUrl": "/media/19782/ltc.png",
					"Name": "LTC",
					"CoinName": "Litecoin",
					"FullName": "Litecoin (LTC)",
					"Algorithm": "Scrypt",
					"ProofType": "PoW",
					"SortOrder": "2"
				}
			},
			"Type": 100
		}`

		fmt.Fprint(w, response)
	})

	acct, _, err := client.Coins.List(ctx)
	if err != nil {
		t.Errorf("Coins.List returned error: %v", err)
	}

	expected := []Coin{{"3808", "/coins/ltc/overview", "/media/19782/ltc.png", "LTC", "Litecoin", "Litecoin (LTC)", "Scrypt", "PoW", "2"}}

	if !reflect.DeepEqual(acct, expected) {
		t.Errorf("Coins.List returned %+v, expected %+v", acct, expected)
	}
}
