package main

import (
	"fmt"

	"github.com/lucazulian/cryptocomparego"
	"github.com/lucazulian/cryptocomparego/context"
)

func main() {

	client := cryptocomparego.NewClient(nil)
	ctx := context.TODO()

	priceMultiRequest := cryptocomparego.NewPriceMultiRequest([]string{"BTC", "ETH"}, []string{"BTC", "USD", "EUR"})
	priceMultiList, _, err := client.PriceMulti.List(ctx, priceMultiRequest)

	if err != nil {
		fmt.Printf("Something bad happened: %s\n", err)
	}

	for _, priceMulti := range priceMultiList {
		for _, coin := range priceMulti.Value {
			fmt.Printf("Main Coin %s, Coin %s - %f\n", priceMulti.Name, coin.Name, coin.Value)
		}
	}
}
