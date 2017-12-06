package cryptocomparego

import (
	"fmt"
	"github.com/lucazulian/cryptocomparego/context"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	pricemultiBasePath = "data/pricemulti"
)

type PriceMultiService interface {
	List(context.Context, *PriceMultiRequest) ([]Price, *Response, error)
}

type PriceMultiServiceOp struct {
	client *Client
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
		values.Add("fsym", strings.Join(pr.Fsyms, ","))
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

func (s *PriceMultiServiceOp) List(ctx context.Context, priceMultiRequest *PriceMultiRequest) ([]Price, *Response, error) {

	path := pricemultiBasePath

	if priceMultiRequest != nil {
		path = priceMultiRequest.FormattedQueryString(priceBasePath)
	}

	_, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	return nil, nil, nil
}
