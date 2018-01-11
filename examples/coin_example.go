package main

import (
	"fmt"

	"github.com/lucazulian/cryptocomparego"
	"github.com/lucazulian/cryptocomparego/context"
)

func main() {

	client := cryptocomparego.NewClient(nil)
	ctx := context.TODO()

	coinList, _, err := client.Coin.List(ctx)

	if err != nil {
		fmt.Printf("Something bad happened: %s\n", err)
	}

	for _, coin := range coinList {
		fmt.Printf("Coin %s - %s\n", coin.Name, coin.FullName)
	}
}
