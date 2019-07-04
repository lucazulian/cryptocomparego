package cryptocomparego

import (
	"net/http"
	"sort"

	"github.com/lucazulian/cryptocomparego/context"
)

const (
	coinBasePath = "data/all/coinlist"
)

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
	Symbol    string `json:"Symbol"`
	CoinName  string `json:"CoinName"`
	FullName  string `json:"FullName"`
	Algorithm string `json:"Algorithm"`
	ProofType string `json:"ProofType"`
	SortOrder string `json:"SortOrder"`
}

//TODO try to remove Sorter duplication
type CoinNamesSorter []Coin

func (a CoinNamesSorter) Len() int           { return len(a) }
func (a CoinNamesSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a CoinNamesSorter) Less(i, j int) bool { return a[i].Name < a[j].Name }

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

	req, err := s.client.NewRequest(ctx, http.MethodGet, *s.client.MinURL, urlPath, nil)
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

	sort.Sort(CoinNamesSorter(coins))

	return coins, resp, err
}
