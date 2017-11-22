package cryptocomparego

import (
	"github.com/lucazulian/cryptocomparego/context"
	"net/http"
)

const coinBasePath = "data/coinlist"

type CoinService interface {
	List(context.Context) ([]Coin, *Response, error)
}

type CoinServiceOp struct {
	client *Client
}

var _ CoinService = &CoinServiceOp{}

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

type coinsRoot struct {
	Response     string          `json:"Response"`
	Message      string          `json:"Message"`
	BaseImageUrl string          `json:"BaseImageUrl"`
	BaseLinkUrl  string          `json:"BaseLinkUrl"`
	Data         map[string]Coin `json:"Data"`
	Type         int             `json:"Type"`
}

func (ds *coinsRoot) GetCoins() ([]Coin, error) {
	var coins []Coin
	for _, value := range ds.Data {
		value.Url = ds.BaseLinkUrl + value.Url
		value.ImageUrl = ds.BaseImageUrl + value.ImageUrl
		coins = append(coins, value)
	}

	return coins, nil
}

func (s *CoinServiceOp) List(ctx context.Context) ([]Coin, *Response, error) {

	urlPath := coinBasePath

	req, err := s.client.NewRequest(ctx, http.MethodGet, urlPath, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(coinsRoot)
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
