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
	pricemultiBasePath = "data/pricemulti"
)

type PriceMultiService interface {
	List(context.Context, *PriceMultiRequest) ([]PriceMulti, *Response, error)
}

type PriceMultiServiceOp struct {
	client *Client
}

type PriceMulti struct {
	Name  string
	Value []Price
}

var _ PriceMultiService = &PriceMultiServiceOp{}

type PriceMultiRequest struct {
	Fsyms         []string
	Tsyms         []string
	E             string
	ExtraParams   string
	Sign          bool
	TryConversion bool
}

func NewPriceMultiRequest(fsyms []string, tsyms []string) *PriceMultiRequest {
	pr := PriceMultiRequest{Fsyms: fsyms, Tsyms: tsyms}
	pr.E = "CCCAGG"
	pr.Sign = false
	pr.TryConversion = true
	return &pr
}

func (pr *PriceMultiRequest) FormattedQueryString(baseUrl string) string {
	values := url.Values{}

	if len(pr.Fsyms) > 0 {
		values.Add("fsyms", strings.Join(pr.Fsyms, ","))
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
type PriceMultiNameSorter []PriceMulti

func (a PriceMultiNameSorter) Len() int           { return len(a) }
func (a PriceMultiNameSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PriceMultiNameSorter) Less(i, j int) bool { return a[i].Name < a[j].Name }

type priceMultiRoot map[string]interface{}

func (ds *priceMultiRoot) GetPrices() ([]PriceMulti, error) {
	var pricesMulti []PriceMulti
	for pKey, pValue := range *ds {
		priceMulti := PriceMulti{}
		priceMulti.Name = pKey
		var prices []Price
		for key, value := range pValue.(map[string]interface{}) {
			price := Price{key, value.(float64)}
			prices = append(prices, price)
		}
		sort.Sort(PriceNameSorter(prices))
		priceMulti.Value = prices
		pricesMulti = append(pricesMulti, priceMulti)
	}

	return pricesMulti, nil
}

func (ds *priceMultiRoot) HasError() error {
	//TODO try to unmarshal with error struct
	var priceMultiError error = nil
	if val, ok := (*ds)["Response"]; ok {
		if val == "Error" {
			val, _ = (*ds)["Message"]
			priceMultiError = errors.New(val.(string))
		}
	}
	return priceMultiError
}

func (s *PriceMultiServiceOp) List(ctx context.Context, priceMultiRequest *PriceMultiRequest) ([]PriceMulti, *Response, error) {

	path := pricemultiBasePath

	if priceMultiRequest != nil {
		path = priceMultiRequest.FormattedQueryString(pricemultiBasePath)
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, *s.client.MinURL, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(priceMultiRoot)
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

	sort.Sort(PriceMultiNameSorter(prices))

	return prices, resp, err
}
