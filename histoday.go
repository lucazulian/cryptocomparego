package cryptocomparego

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/pkg/errors"

	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"time"

	"github.com/lucazulian/cryptocomparego/context"
)

const (
	histodyBasePath = "data/histoday"
)

// Get the history kline data of any cryptocurrency in any other currency that you need.
type HistodayService interface {
	Get(context.Context, *HistodayRequest) ([]Histoday, *Response, error)
}

type HistodayServiceOp struct {
	client *Client
}

var _ HistodayService = &HistodayServiceOp{}

type conversionType struct {
	Type             string `json:"type"`
	ConversionSymbol string `json:"conversionSymbol"`
}

type histodayResp struct {
	Response          string         `json:"Response"`
	Message           string         `json:"Message"` // Error Message
	Type              int            `json:"Type"`
	Aggregated        bool           `json:"Aggregated"`
	Data              []Histoday     `json:"Data"`
	TimeTo            int64          `json:"TimeTo"`
	TimeFrom          int64          `json:"TimeFrom"`
	FirstValueInArray bool           `json:"FirstValueInArray"`
	ConversionType    conversionType `json:"ConversionType"`
}

type Histoday struct {
	Time       int64   `json:"time"`
	Close      float64 `json:"close"`
	High       float64 `json:"high"`
	Low        float64 `json:"low"`
	Open       float64 `json:"open"`
	VolumeFrom float64 `json:"volumefrom"`
	VolumeTo   float64 `json:"volumeto"`
}

type HistodayRequest struct {
	Fsym          string
	Tsym          string
	E             string
	ExtraParams   string
	Sign          bool
	TryConversion bool
	Aggregate     int // Not Used For Now
	Limit         int
	ToTs          time.Time // Not Used For Now
	AllData       bool
}

func NewHistodayRequest(fsym string, tsym string, limitDays int, allData bool) *HistodayRequest {
	pr := HistodayRequest{Fsym: fsym, Tsym: tsym}
	pr.E = "CCCAGG"
	pr.Sign = false
	pr.TryConversion = true
	pr.Aggregate = 1
	if limitDays < 1 {
		limitDays = 1
	}
	if limitDays > 2000 {
		limitDays = 2000
	}
	pr.Limit = limitDays
	pr.AllData = allData
	return &pr
}

func (hr *HistodayRequest) FormattedQueryString(baseUrl string) string {
	values := url.Values{}

	if len(hr.Fsym) > 0 {
		values.Add("fsym", hr.Fsym)
	}

	if len(hr.Tsym) > 0 {
		values.Add("tsym", hr.Tsym)
	}

	if len(hr.E) > 0 {
		values.Add("e", hr.E)
	}

	if len(hr.ExtraParams) > 0 {
		values.Add("extraParams", hr.ExtraParams)
	}

	values.Add("sign", strconv.FormatBool(hr.Sign))
	values.Add("tryConversion", strconv.FormatBool(hr.TryConversion))
	values.Add("limit", strconv.FormatInt(int64(hr.Limit), 10))
	values.Add("allData", strconv.FormatBool(hr.AllData))

	return fmt.Sprintf("%s?%s", baseUrl, values.Encode())
}

func ReadAndAssignResponseBody(res *http.Response) (io.Reader, error) {
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	res.Body = ioutil.NopCloser(bytes.NewReader(buf))
	return bytes.NewReader(buf), nil
}

func (s *HistodayServiceOp) Get(ctx context.Context, histodayRequest *HistodayRequest) ([]Histoday, *Response, error) {

	path := histodyBasePath

	if histodayRequest != nil {
		path = histodayRequest.FormattedQueryString(histodyBasePath)
	}

	reqUrl := fmt.Sprintf("%s%s", s.client.MinURL.String(), path)
	resp, err := http.Get(reqUrl)
	res := Response{}
	res.Response = resp
	if err != nil {
		return nil, &res, err
	}
	defer func() { resp.Body.Close() }()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, &res, err
	}
	if len(buf) <= 0 {
		return nil, &res, errors.New("Empty response")
	}

	hr := histodayResp{}
	err = json.Unmarshal(buf, &hr)
	if err != nil {
		return nil, &res, errors.Wrap(err, fmt.Sprintf("JSON Unmarshal error, raw string is '%s'", string(buf)))
	}
	if hr.Response == "Error" {
		return nil, &res, errors.New(hr.Message)
	}

	return hr.Data, &res, nil
}
