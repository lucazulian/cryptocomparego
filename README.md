[![Build Status](https://api.travis-ci.org/lucazulian/cryptocomparego.svg)](https://travis-ci.org/lucazulian/cryptocomparego) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

# Cryptocomparego

Cryptocomparego is a Go client library for accessing the Cryptocompare API.

You can view the client API docs here: [http://godoc.org/github.com/lucazulian/cryptocomparego](http://godoc.org/github.com/lucazulian/cryptocomparego)

You can view Cryptocompare API docs here: [https://www.cryptocompare.com/api/](https://www.cryptocompare.com/api/)


## Usage

```go
import "github.com/lucazulian/cryptocomparego"
```

## Examples


To get general info for all the coins available:

```go
ctx := context.TODO()

coinList, _, err := client.Coin.List(ctx)

if err != nil {
    fmt.Printf("Something bad happened: %s\n", err)
    return err
}
```

## Contributing

We love pull requests! Please see the [contribution guidelines](CONTRIBUTING.md).