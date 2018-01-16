package cryptocomparego

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"

	"github.com/lucazulian/cryptocomparego/context"
)

const (
	priceBasePath = "data/price"
)

// Get the current price of any cryptocurrency in any other currency that you need.
// If the crypto does not trade directly into the toSymbol requested, BTC will be used for conversion.
// If the oposite pair trades we invert it (eg.: BTC-XMR).
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

func NewPriceRequest(fsym string, tsyms []string) *PriceRequest {
	pr := PriceRequest{Fsym: fsym, Tsyms: tsyms}
	pr.E = "CCCAGG"
	pr.Sign = false
	pr.TryConversion = true
	return &pr
}

func (pr *PriceRequest) FormattedQueryString(baseUrl string) string {
	values := url.Values{}

	if len(pr.Fsym) > 0 {
		values.Add("fsym", pr.Fsym)
	}

	if len(pr.Tsyms) > 0 {
		values.Add("tsyms", strings.Join(pr.Tsyms, ","))
	}

	if len(pr.E) > 0 {
		values.Add("e", pr.E)
	}

	if len(pr.ExtraParams) > 0 {
		values.Add("extraParams", pr.ExtraParams)
	}

	values.Add("sign", strconv.FormatBool(pr.Sign))
	values.Add("tryConversion", strconv.FormatBool(pr.TryConversion))

	return fmt.Sprintf("%s?%s", baseUrl, values.Encode())
}

//TODO try to remove Sorter duplication
type PriceNameSorter []Price

func (a PriceNameSorter) Len() int           { return len(a) }
func (a PriceNameSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PriceNameSorter) Less(i, j int) bool { return a[i].Name < a[j].Name }

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
		path = priceRequest.FormattedQueryString(priceBasePath)
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, *s.client.MinURL, path, nil)
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

	sort.Sort(PriceNameSorter(prices))

	return prices, resp, err
}
