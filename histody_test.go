package cryptocomparego

import (
	"testing"
)

func TestFormattedQueryStringNilHistodayRequest(t *testing.T) {

	histodayRequest := NewHistodayRequest("", "", 10, false)
	acct := histodayRequest.FormattedQueryString("/data/histoday")

	expected := "/data/histoday?allData=false&e=CCCAGG&limit=10&sign=false&tryConversion=true"

	if acct != expected {
		t.Errorf("HistodayRequest.FormattedQueryString returned %+v, expected %+v", acct, expected)
	}
}

func TestFormattedQueryStringHistodayRequest(t *testing.T) {

	histodayRequest := NewHistodayRequest("ETH", "BTC", 30, false)
	acct := histodayRequest.FormattedQueryString("/data/histoday")

	expected := "/data/histoday?allData=false&e=CCCAGG&fsym=ETH&limit=30&sign=false&tryConversion=true&tsym=BTC"

	if acct != expected {
		t.Errorf("HistodayRequest.FormattedQueryString returned %+v, expected %+v", acct, expected)
	}
}
