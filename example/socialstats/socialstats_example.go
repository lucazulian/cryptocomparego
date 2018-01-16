package main

import (
	"fmt"

	"github.com/lucazulian/cryptocomparego"
	"github.com/lucazulian/cryptocomparego/context"
)

func main() {

	client := cryptocomparego.NewClient(nil)
	ctx := context.TODO()

	socialStats, _, err := client.SocialStats.Get(ctx, "1182")

	if err != nil {
		fmt.Printf("Something bad happened: %s\n", err)
	}

	fmt.Printf("Stats %+v\n", socialStats)
}
