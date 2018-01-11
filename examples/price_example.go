package main

import (
	"fmt"

	"github.com/lucazulian/cryptocomparego"
	"github.com/lucazulian/cryptocomparego/context"
)

func main() {

	client := cryptocomparego.NewClient(nil)
	ctx := context.TODO()

	priceRequest := cryptocomparego.NewPriceRequest("ETH", []string{"BTC", "USD", "EUR"})
	priceList, _, err := client.Price.List(ctx, priceRequest)

	if err != nil {
		fmt.Printf("Something bad happened: %s\n", err)
	}

	for _, coin := range priceList {
		fmt.Printf("Coin %s - %f\n", coin.Name, coin.Value)
	}
}
