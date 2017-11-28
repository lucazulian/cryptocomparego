package cryptocomparego

import (
	"errors"
	"fmt"
	"github.com/lucazulian/cryptocomparego/context"
	"net/http"
	"sort"
	"strings"
)

const priceBasePath = "data/price"

type PriceService interface {
	List(context.Context, *PriceRequest) ([]Price, *Response, error)
}

type PriceServiceOp struct {
	client *Client
}

var _ PriceService = &PriceServiceOp{}

type Price struct {
	Name  string
	Value float64
}

type PriceRequest struct {
	Fsym          string
	Tsyms         []string
	E             string
	ExtraParams   string
	Sign          bool
	TryConversion bool
}

type PriceNamesSorter []Price

func (a PriceNamesSorter) Len() int           { return len(a) }
func (a PriceNamesSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PriceNamesSorter) Less(i, j int) bool { return a[i].Name < a[j].Name }

//type priceRoot map[string]float64
type priceRoot map[string]interface{}

func (ds *priceRoot) GetPrices() ([]Price, error) {
	var prices []Price
	for key, value := range *ds {
		price := Price{key, value.(float64)}
		prices = append(prices, price)
	}

	return prices, nil
}

func (ds *priceRoot) HasError() error {
	//TODO try to unmarshal with error struct
	var priceError error = nil
	if val, ok := (*ds)["Response"]; ok {
		if val == "Error" {
			val, _ = (*ds)["Message"]
			priceError = errors.New(val.(string))
		}
	}
	return priceError
}

func (s *PriceServiceOp) List(ctx context.Context, priceRequest *PriceRequest) ([]Price, *Response, error) {

	path := priceBasePath

	if priceRequest != nil {
		path = fmt.Sprintf("%s?fsym=%s&tsyms=%s", priceBasePath, priceRequest.Fsym, strings.Join(priceRequest.Tsyms, ","))
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(priceRoot)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	if err := root.HasError(); err != nil {
		return nil, resp, err
	}

	prices, err := root.GetPrices()
	if err != nil {
		return nil, resp, err
	}

	sort.Sort(PriceNamesSorter(prices))

	return prices, resp, err
}
