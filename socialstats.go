package cryptocomparego

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/lucazulian/cryptocomparego/context"
)

const (
	socialstatsBasePath = "data/socialstats"
)

type SocialStatsService interface {
	Get(context.Context, int) (*SocialStats, *Response, error)
}

type SocialStatsServiceOp struct {
	client *Client
}

var _ SocialStatsService = &SocialStatsServiceOp{}

type General struct {
	Name     string      `json:"Name"`
	CoinName interface{} `json:"CoinName"`
	Type     interface{} `json:"Type"`
	Points   int         `json:"Points"`
}

type PageViewsSplit struct {
	Overview  int `json:"Overview"`
	Markets   int `json:"Markets"`
	Analysis  int `json:"Analysis"`
	Charts    int `json:"Charts"`
	Trades    int `json:"Trades"`
	Forum     int `json:"Forum"`
	Influence int `json:"Influence"`
}

type SimilarItem struct {
	Id            int    `json:"Id"`
	Name          string `json:"Name"`
	FullName      string `json:"FullName"`
	ImageUrl      string `json:"ImageUrl"`
	Url           string `json:"Url"`
	FollowingType int    `json:"FollowingType"`
}

type CryptopianFollower struct {
	Id       int    `json:"Id"`
	Name     string `json:"Name"`
	ImageUrl string `json:"ImageUrl"`
	Url      string `json:"Url"`
	Type     string `json:"Type"`
}

type CryptoCompare struct {
	SimilarItems        []SimilarItem        `json:"SimilarItems"`
	CryptopianFollowers []CryptopianFollower `json:"CryptopianFollowers"`
	Comments            string               `json:"Comments"`
	Points              int                  `json:"Points"`
	Posts               string               `json:"Posts"`
	Followers           int                  `json:"Followers"`
	PageViewsSplit      PageViewsSplit       `json:"PageViewsSplit"`
	PageViews           int                  `json:"PageViews"`
}

type Twitter struct {
	Followers       int    `json:"followers"`
	Following       string `json:"following"`
	Lists           int    `json:"lists"`
	Favourites      string `json:"favourites"`
	Statuses        int    `json:"statuses"`
	AccountCreation string `json:"account_creation"`
	Name            string `json:"name"`
	Link            string `json:"link"`
	Points          int    `json:"Points"`
}

type Reddit struct {
	Subscribers       int     `json:"subscribers"`
	ActiveUsers       int     `json:"active_users"`
	CommunityCreation string  `json:"community_creation"`
	PostsPerHour      float64 `json:"posts_per_hour,string"`
	PostsPerDay       float64 `json:"posts_per_day,string"`
	CommentsPerHour   float64 `json:"comments_per_hour,string"`
	CommentsPerDay    float64 `json:"comments_per_day"`
	Link              string  `json:"link"`
	Name              string  `json:"name"`
	Points            int     `json:"Points"`
}

type Facebook struct {
	Likes        int    `json:"likes"`
	IsClosed     bool   `json:"is_closed,string"`
	TalkingAbout string `json:"talking_about"`
	Name         string `json:"name"`
	Link         string `json:"link"`
	Points       int    `json:"Points"`
}

type Source struct {
	Name       string `json:"Name"`
	Url        string `json:"Url"`
	InternalId int    `json:"InternalId"`
}

type CodeRepositoryItem struct {
	Stars             int    `json:"stars"`
	Language          string `json:"language"`
	Forks             int    `json:"forks"`
	OpenTotalIssues   string `json:"open_total_issues"`
	Subscribers       int    `json:"subscribers"`
	Size              string `json:"size"`
	Url               string `json:"url"`
	LastUpdate        string `json:"last_update"`
	LastPush          string `json:"last_push"`
	CreatedAt         string `json:"created_at"`
	Fork              bool   `json:"fork,string"`
	Source            Source `json:"source"`
	Parent            Source `json:"parent"`
	OpenPullIssues    string `json:"open_pull_issues"`
	ClosedPullIssues  string `json:"closed_pull_issues"`
	ClosedTotalIssues string `json:"closed_total_issues"`
	OpenIssues        string `json:"open_issues"`
	ClosedIssues      string `json:"closed_issues"`
}

type CodeRepository struct {
	List   []CodeRepositoryItem `json:"List"`
	Points int                  `json:"Points"`
}

type SocialStats struct {
	General        General        `json:"General"`
	CryptoCompare  CryptoCompare  `json:"CryptoCompare"`
	Twitter        Twitter        `json:"Twitter"`
	Reddit         Reddit         `json:"Reddit"`
	Facebook       Facebook       `json:"Facebook"`
	CodeRepository CodeRepository `json:"CodeRepository"`
}

type socialStatsRoot struct {
	Response     string      `json:"Response"`
	Message      string      `json:"Message"`
	BaseImageUrl string      `json:"BaseImageUrl"`
	BaseLinkUrl  string      `json:"BaseLinkUrl"`
	Data         SocialStats `json:"Data"`
	Type         int         `json:"Type"`
}

func (s *SocialStatsServiceOp) Get(ctx context.Context, coinExchangeId int) (*SocialStats, *Response, error) {

	values := url.Values{}
	values.Add("id", strconv.Itoa(coinExchangeId))

	path := fmt.Sprintf("%s?%s", socialstatsBasePath, values.Encode())

	req, err := s.client.NewRequest(ctx, http.MethodGet, *s.client.BaseURL, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(socialStatsRoot)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return &root.Data, resp, err
}
