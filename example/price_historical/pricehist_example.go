package main

import (
	"fmt"

	"github.com/lucazulian/cryptocomparego/context"
	"github.com/nwops/cryptocomparego"
)

func main() {

	client := cryptocomparego.NewClient(nil)
	ctx := context.TODO()

	// Using a 0 as the timestamp gets the current time
	// You can use time.Now().Unix() to get the current time and manipulate it from there
	// Or pass in a unix timestamp that was computed elsewhere
	//priceHistRequest := NewPriceHistRequest("ETH", []string{"BTC", "USD", "EUR"}, 1560918228)
	priceHistRequest := cryptocomparego.NewPriceHistRequest("ETH", []string{"BTC", "USD", "EUR"}, 0)

	priceHistList, _, err := client.PriceHist.List(ctx, priceHistRequest)

	if err != nil {
		fmt.Printf("Something bad happened: %s\n", err)
	}

	for _, coin := range priceHistList {
		fmt.Printf("Coin %s - %f\n", coin.Name, coin.Value)
	}
}
