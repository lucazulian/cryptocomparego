package cryptocomparego

import (
	"testing"
)

func TestFormattedQueryStringNilHistominuteRequest(t *testing.T) {

	histominuteRequest := NewHistominuteRequest("", "", 10, -1)
	acct := histominuteRequest.FormattedQueryString("/data/histominute")

	expected := "/data/histominute?e=CCCAGG&limit=10&sign=false&tryConversion=true"

	if acct != expected {
		t.Errorf("HistodayRequest.FormattedQueryString returned %+v, expected %+v", acct, expected)
	}
}

func TestFormattedQueryStringHistominuteRequest(t *testing.T) {

	histodayRequest := NewHistominuteRequest("ETH", "BTC", 30, 10)
	acct := histodayRequest.FormattedQueryString("/data/histominute")

	expected := "/data/histominute?e=CCCAGG&fsym=ETH&limit=30&sign=false&toTs=10&tryConversion=true&tsym=BTC"

	if acct != expected {
		t.Errorf("HistodayRequest.FormattedQueryString returned %+v, expected %+v", acct, expected)
	}
}
