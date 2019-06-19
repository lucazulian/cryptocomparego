package cryptocomparego

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
	"github.com/lucazulian/cryptocomparego/context"
)

const (
	libraryVersion = "0.1.0"
	defaultBaseURL = "https://www.cryptocompare.com/api/"
	minBaseURL     = "https://min-api.cryptocompare.com/"
	userAgent      = "cryptocomparego/" + libraryVersion
	mediaType      = "application/json"
)

type Client struct {
	client *http.Client

	BaseURL *url.URL

	MinURL *url.URL

	UserAgent string

	Coin CoinService

	Price PriceService

	PriceMulti PriceMultiService

	PriceMultiFull PriceMultiFullService

	SocialStats SocialStatsService

	Histoday HistodayService

	Histohour HistohourService

	Histomin HistominuteService

	onRequestCompleted RequestCompletionCallback
}

type RequestCompletionCallback func(*http.Request, *http.Response)

type Response struct {
	*http.Response

	Monitor string
}

type ErrorResponse struct {
	Response  *http.Response
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
}

func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)

	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	origURL, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	origValues := origURL.Query()

	newValues, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	for k, v := range newValues {
		origValues[k] = v
	}

	origURL.RawQuery = origValues.Encode()
	return origURL.String(), nil
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)
	minURL, _ := url.Parse(minBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL, MinURL: minURL, UserAgent: userAgent}
	c.Coin = &CoinServiceOp{client: c}
	c.Price = &PriceServiceOp{client: c}
	c.PriceMulti = &PriceMultiServiceOp{client: c}
	c.PriceMultiFull = &PriceMultiFullServiceOp{client: c}

	c.SocialStats = &SocialStatsServiceOp{client: c}
	c.Histoday = &HistodayServiceOp{client: c}
	c.Histomin = &HistominuteServiceOp{client: c}
	c.Histohour = &HistohourServiceOp{client: c}

	return c
}

type ClientOpt func(*Client) error

func New(httpClient *http.Client, opts ...ClientOpt) (*Client, error) {
	c := NewClient(httpClient)
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func SetBaseURL(bu string) ClientOpt {
	return func(c *Client) error {
		u, err := url.Parse(bu)
		if err != nil {
			return err
		}

		c.BaseURL = u
		return nil
	}
}

func SetUserAgent(ua string) ClientOpt {
	return func(c *Client) error {
		c.UserAgent = fmt.Sprintf("%s %s", ua, c.UserAgent)
		return nil
	}
}

func (c *Client) NewRequest(ctx context.Context, method string, baseUrl url.URL, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := baseUrl.ResolveReference(rel)

	buf := new(bytes.Buffer)
	if body != nil {
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", mediaType)
	req.Header.Add("Accept", mediaType)
	req.Header.Add("User-Agent", c.UserAgent)
	return req, nil
}

func (c *Client) OnRequestCompleted(rc RequestCompletionCallback) {
	c.onRequestCompleted = rc
}

func newResponse(r *http.Response) *Response {
	response := Response{Response: r}
	return &response
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	resp, err := context.DoRequestWithClient(ctx, c.client, req)

	if err != nil {
		return nil, err
	}
	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, resp)
	}

	defer func() {
		if bErr := resp.Body.Close(); err == nil {
			err = bErr
		}
	}()

	response := newResponse(resp)

	err = CheckResponse(resp)

	if err != nil {
		return nil, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
			if err != nil {
				return nil, err
			}
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err != nil {
				return nil, err
			}
		}
	}

	return response, nil
}

func (r *ErrorResponse) Error() string {
	if r.RequestID != "" {
		return fmt.Sprintf("%v %v: %d (request %q) %v",
			r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.RequestID, r.Message)
	}
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Message)
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && len(data) > 0 {
		err := json.Unmarshal(data, errorResponse)
		if err != nil {
			return err
		}
	}

	return errorResponse
}
