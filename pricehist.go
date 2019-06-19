package cryptocomparego

import (
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/lucazulian/cryptocomparego/context"
)

const (
	priceHistBasePath = "data/pricehistorical"
)

// Get the current price of any cryptocurrency in any other currency that you need.
// If the crypto does not trade directly into the toSymbol requested, BTC will be used for conversion.
// If the oposite pair trades we invert it (eg.: BTC-XMR).
type PriceHistService interface {
	List(context.Context, *PriceHistRequest) ([]PriceHist, *Response, error)
}

type PriceHist struct {
	Name  string
	Value float64
}

type PriceHistServiceOp struct {
	client *Client
}
type PriceHistNameSorter []PriceHist

var _ PriceHistService = &PriceHistServiceOp{}

type PriceHistRequest struct {
	Fsym            string
	Tsyms           []string
	E               string
	ExtraParams     string
	Ts              int64
	CalculationType string
	Sign            bool
	TryConversion   bool
}

func NewPriceHistRequest(fsym string, tsyms []string, ts int64) *PriceHistRequest {
	if ts < 1 {
		ts = time.Now().Unix()
	}
	pr := PriceHistRequest{Fsym: fsym, Tsyms: tsyms}
	pr.E = "CCCAGG"
	pr.Ts = ts
	pr.CalculationType = "Close"
	pr.Sign = false
	pr.TryConversion = true
	return &pr
}

func (pr *PriceHistRequest) FormattedQueryString(baseUrl string) string {
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

	if pr.Ts > 0 {
		values.Add("ts", strconv.FormatInt(pr.Ts, 10))
	}

	if len(pr.CalculationType) > 0 {
		values.Add("calculationType", pr.CalculationType)
	}

	if len(pr.ExtraParams) > 0 {
		values.Add("extraParams", pr.ExtraParams)
	}

	values.Add("sign", strconv.FormatBool(pr.Sign))
	values.Add("tryConversion", strconv.FormatBool(pr.TryConversion))

	return fmt.Sprintf("%s?%s", baseUrl, values.Encode())
}
func (a PriceHistNameSorter) Len() int           { return len(a) }
func (a PriceHistNameSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PriceHistNameSorter) Less(i, j int) bool { return a[i].Name < a[j].Name }

func (ds *priceRoot) GetHistPrices() ([]PriceHist, error) {
	var prices []PriceHist
	for _, value := range *ds {
		for coinPair, coinPairPrice := range value.(map[string]interface{}) {
			priceHist := PriceHist{coinPair, coinPairPrice.(float64)}
			prices = append(prices, priceHist)
		}
	}

	return prices, nil
}

func (s *PriceHistServiceOp) List(ctx context.Context, priceHistRequest *PriceHistRequest) ([]PriceHist, *Response, error) {

	path := priceHistBasePath

	if priceHistRequest != nil {
		path = priceHistRequest.FormattedQueryString(priceHistBasePath)
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

	prices, err := root.GetHistPrices()
	if err != nil {
		return nil, resp, err
	}

	sort.Sort(PriceHistNameSorter(prices))

	return prices, resp, err
}
