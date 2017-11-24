package cryptocomparego

import (
	"github.com/lucazulian/cryptocomparego/context"
	"net/http"
)

const priceBasePath = "data/price"

type PriceService interface {
	List(context.Context) ([]Price, *Response, error)
}

type PriceServiceOp struct {
	client *Client
}

var _ PriceService = &PriceServiceOp{}

type Price struct {
	Name  string
	Value float64
}

type priceRoot map[string]float64

func (ds *priceRoot) GetPrices() ([]Price, error) {
	var prices []Price
	for key, value := range *ds {
		price := Price{key, value}
		prices = append(prices, price)
	}

	return prices, nil
}

func (s *PriceServiceOp) List(ctx context.Context) ([]Price, *Response, error) {

	path := priceBasePath
	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(priceRoot)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	prices, err := root.GetPrices()
	if err != nil {
		return nil, resp, err
	}
	return prices, resp, err
}
