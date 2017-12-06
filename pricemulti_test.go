package cryptocomparego

import "testing"

func TestFormattedQueryStringNilPriceMultiRequest(t *testing.T) {

	priceRequest := NewPriceMultiRequest([]string{}, nil)
	acct := priceRequest.FormattedQueryString("/data/pricemulti")

	expected := "/data/pricemulti?e=CCCAGG&sign=false&tryConversion=true"

	if acct != expected {
		t.Errorf("PriceRequest.FormattedQueryString returned %+v, expected %+v", acct, expected)
	}
}
