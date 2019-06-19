package cryptocomparego

import (
	"testing"
	"time"
)

func TestFormattedQueryStringNilHistohourRequest(t *testing.T) {

	histohourRequest := NewHistohourRequest("", "", 10, time.Time{})
	acct := histohourRequest.FormattedQueryString("/data/histohour")

	expected := "/data/histohour?e=CCCAGG&limit=10&sign=false&tryConversion=true"

	if acct != expected {
		t.Errorf("HistodayRequest.FormattedQueryString returned %+v, expected %+v", acct, expected)
	}
}

func TestFormattedQueryStringHistohourRequest(t *testing.T) {

	histodayRequest := NewHistohourRequest("ETH", "BTC", 30, time.Unix(10, 0))
	acct := histodayRequest.FormattedQueryString("/data/histohour")

	expected := "/data/histohour?e=CCCAGG&fsym=ETH&limit=30&sign=false&toTs=10&tryConversion=true&tsym=BTC"

	if acct != expected {
		t.Errorf("HistodayRequest.FormattedQueryString returned %+v, expected %+v", acct, expected)
	}
}
