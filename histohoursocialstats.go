package cryptocomparego

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/lucazulian/cryptocomparego/context"
)

const (
	histohoursocialstatsBasePath = "data/social/coin/histo/hour"
)

type HistohourSocialStatsService interface {
	Get(context.Context, *HistohourSocialStatsRequest) (*HistohourSocialStatsResponse, *Response, error)
}

type HistohourSocialStatsServiceOp struct {
	client *Client
}

var _ HistohourSocialStatsService = &HistohourSocialStatsServiceOp{}

type HistohourSocialStatsResponse struct {
	Response   string                 `json:"Response"`
	Message    string                 `json:"Message"` // Error Message
	HasWarning bool                   `json:"HasWarning"`
	Type       int                    `json:"Type"`
	Data       []HistohourSocialStats `json:"Data,omitempty"`
}

type HistohourSocialStats struct {
	Time                     int64   `json:"time"`
	Comments                 int64   `json:"comments"`
	Posts                    int64   `json:"posts"`
	Followers                int64   `json:"followers"`
	Points                   int64   `json:"points"`
	OverviewPageViews        int64   `json:"overview_page_views"`
	AnalysisPageViews        int64   `json:"analysis_page_views"`
	MarketsPageViews         int64   `json:"markets_page_views"`
	ChartsPageViews          int64   `json:"charts_page_views"`
	TradesPageViews          int64   `json:"trades_page_views"`
	ForumPageViews           int64   `json:"forum_page_views"`
	InfluencePageViews       int64   `json:"influence_page_views"`
	TotalPageViews           int64   `json:"total_page_views"`
	FbLikes                  int64   `json:"fb_likes"`
	FbTalkingAbout           int64   `json:"fb_talking_about"`
	TwitterFollowers         int64   `json:"twitter_followers"`
	TwitterFollowing         int64   `json:"twitter_following"`
	TwitterLists             int64   `json:"twitter_lists"`
	TwitterFavourites        int64   `json:"twitter_favourites"`
	TwitterStatuses          int64   `json:"twitter_statuses"`
	RedditSubscribers        int64   `json:"reddit_subscribers"`
	RedditActiveUsers        int64   `json:"reddit_active_users"`
	RedditPostsPerHour       float64 `json:"reddit_posts_per_hour"`
	RedditPostsPerDay        float64 `json:"reddit_posts_per_day"`
	RedditCommentsPerHour    float64 `json:"reddit_comments_per_hour"`
	RedditCommentsPerDay     float64 `json:"reddit_comments_per_day"`
	CodeRepoStars            int64   `json:"code_repo_stars"`
	CodeRepoForks            int64   `json:"code_repo_forks"`
	CodeRepoSubscribers      int64   `json:"code_repo_subscribers"`
	CodeRepoOpenPullIssues   int64   `json:"code_repo_open_pull_issues"`
	CodeRepoClosedPullIssues int64   `json:"code_repo_closed_pull_issues"`
	CodeRepoOpenIssues       int64   `json:"code_repo_open_issues"`
	CodeRepoClosedIssues     int64   `json:"code_repo_closed_issues"`
}

type HistohourSocialStatsRequest struct {
	CoinId      int
	Limit       int
	ToTs        time.Time
	ExtraParams string
	Sign        bool

	Aggregate                       int // Not Used For Now
	AggregatePredictableTimePeriods bool
}

func NewHistohourSocialStatsRequest(coinId int, limit int, toTs time.Time) *HistohourSocialStatsRequest {
	pr := HistohourSocialStatsRequest{
		CoinId:    coinId,
		Limit:     limit,
		ToTs:      toTs,
		Sign:      false,
		Aggregate: 1,
	}

	if limit < 1 {
		limit = 1
	}
	if limit > 2000 {
		limit = 2000
	}
	pr.Limit = limit
	return &pr
}

func (hr *HistohourSocialStatsRequest) FormattedQueryString(baseUrl string) string {
	values := url.Values{}

	values.Add("coinId", strconv.FormatInt(int64(hr.CoinId), 10))
	if len(hr.ExtraParams) > 0 {
		values.Add("extraParams", hr.ExtraParams)
	}

	values.Add("sign", strconv.FormatBool(hr.Sign))
	values.Add("limit", strconv.FormatInt(int64(hr.Limit), 10))
	if hr.ToTs.Unix() >= 0 {
		values.Add("toTs", strconv.FormatInt(int64(hr.ToTs.Unix()), 10))
	}
	values.Add("aggregatePredictableTimePeriods", strconv.FormatBool(hr.AggregatePredictableTimePeriods))

	return fmt.Sprintf("%s?%s", baseUrl, values.Encode())
}

func (s *HistohourSocialStatsServiceOp) Get(ctx context.Context, req *HistohourSocialStatsRequest) (*HistohourSocialStatsResponse, *Response, error) {
	path := histohoursocialstatsBasePath

	if req != nil {
		path = req.FormattedQueryString(histohoursocialstatsBasePath)
	}

	httpReq, err := s.client.NewRequest(ctx, http.MethodGet, *s.client.MinURL, path, nil)
	if err != nil {
		return nil, nil, err
	}

	hr := HistohourSocialStatsResponse{}
	resp, err := s.client.Do(ctx, httpReq, &hr)
	if err != nil {
		return nil, resp, err
	}

	return &hr, resp, nil
}
