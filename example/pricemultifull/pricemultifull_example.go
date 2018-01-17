package main

import (
	"fmt"

	"github.com/lucazulian/cryptocomparego"
	"github.com/lucazulian/cryptocomparego/context"
)

func main() {

	client := cryptocomparego.NewClient(nil)
	ctx := context.TODO()

	priceMultiFullRequest := cryptocomparego.NewPriceMultiFullRequest([]string{"ETH", "DASH"}, []string{"BTC", "USD", "EUR"})
	//priceMultiFullList, _, err := client.PriceMultiFull.List(ctx, priceMultiFullRequest)
	_, _, err := client.PriceMultiFull.List(ctx, priceMultiFullRequest)

	if err != nil {
		fmt.Printf("Something bad happened: %s\n", err)
	}

	//for _, priceMulti := range priceMultiFullList {
	//	for _, coin := range priceMulti.Value {
	//		fmt.Printf("Main Coin %s, Coin %s - %f\n", priceMulti.Name, coin.Name, coin.Value)
	//	}
	//}
}
