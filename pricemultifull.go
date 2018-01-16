package cryptocomparego

import (
	"net/http"

	"github.com/lucazulian/cryptocomparego/context"
)

const (
	pricemultifullBasePath = "data/pricemultifull"
)

// Get all the current trading info (price, vol, open, high, low etc) of any list of cryptocurrencies in any other currency that you need.
// If the crypto does not trade directly into the toSymbol requested, BTC will be used for conversion.
// This API also returns Display values for all the fields.If the oposite pair trades we invert it (eg.: BTC-XMR)".
type PriceMultiFullService interface {
	List(context.Context, *PriceMultiRequest) ([]PriceMultiFull, *Response, error)
}

type PriceMultiFullServiceOp struct {
	client *Client
}

type PriceFull struct {
	Type string `json:"TYPE"`
}

type PriceMultiFullAgg struct {
	Agg map[string]PriceFull
}

type PriceMultiFull struct {
	Raw     map[string]map[string]PriceMultiFullAgg `json:"RAW"`
	Display map[string]map[string]PriceMultiFullAgg `json:"DISPLAY"`
}

var _ PriceMultiFullService = &PriceMultiFullServiceOp{}

func NewPriceMultiFullRequest(fsyms []string, tsyms []string) *PriceMultiRequest {
	pr := PriceMultiRequest{Fsyms: fsyms, Tsyms: tsyms}
	pr.E = "CCCAGG"
	pr.Sign = false
	pr.TryConversion = true
	return &pr
}

func (s *PriceMultiFullServiceOp) List(ctx context.Context, priceMultiRequest *PriceMultiRequest) ([]PriceMultiFull, *Response, error) {

	path := pricemultifullBasePath

	if priceMultiRequest != nil {
		path = priceMultiRequest.FormattedQueryString(pricemultiBasePath)
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, *s.client.MinURL, path, nil)
	if err != nil {
		return nil, nil, err
	}

	_ = req

	return nil, nil, nil
}
