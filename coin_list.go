package cryptocomparego

import (
	"github.com/lucazulian/cryptocomparego/context"
	"net/http"
)

const coinlistBasePath = "data/coinlist"

type CoinListService interface {
	List(context.Context) ([]Coin, *Response, error)
}

type CoinListServiceOp struct {
	client *Client
}

var _ CoinListService = &CoinListServiceOp{}

type Coin struct {
	Id        string `json:"Id"`
	Url       string `json:"Url"`
	ImageUrl  string `json:"ImageUrl"`
	Name      string `json:"Name"`
	CoinName  string `json:"CoinName"`
	FullName  string `json:"FullName"`
	Algorithm string `json:"Algorithm"`
	ProofType string `json:"ProofType"`
	SortOrder string `json:"SortOrder"`
}

type domainsRoot struct {
	Response     string          `json:"Response"`
	Message      string          `json:"Message"`
	BaseImageUrl string          `json:"BaseImageUrl"`
	BaseLinkUrl  string          `json:"BaseLinkUrl"`
	Data         map[string]Coin `json:"Data"`
	Type         int             `json:"Type"`
}

func (ds *domainsRoot) GetCoins() ([]Coin, error) {
	var values []Coin
	for _, value := range ds.Data {
		value.Url = ds.BaseLinkUrl + value.Url
		value.ImageUrl = ds.BaseImageUrl + value.ImageUrl
		values = append(values, value)
	}

	return values, nil
}

func (s *CoinListServiceOp) List(ctx context.Context) ([]Coin, *Response, error) {

	urlPath := coinlistBasePath

	req, err := s.client.NewRequest(ctx, http.MethodGet, urlPath, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(domainsRoot)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	coins, err := root.GetCoins()
	if err != nil {
		return nil, resp, err
	}

	return coins, resp, err
}
