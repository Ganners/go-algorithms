package simplest_url

import (
	"reflect"
	"testing"
)

func TestFindWhitelistQueryParams(t *testing.T) {
	for _, test := range []struct {
		URLs                 []Response
		WhitelistQueryParams []string
	}{
		{
			URLs: []Response{
				{"https://www.buzzfeed.com/foo?id=1&page=1&foo=bar", "1"},
				{"https://www.buzzfeed.com/foo?id=1&page=1&bar=baz", "1"},
				{"https://www.buzzfeed.com/foo?id=1&page=2&bar=baz", "2"},
				{"https://www.buzzfeed.com/foo?id=2&page=1&bing=boo&bar=baz", "3"},
				{"https://www.buzzfeed.com/foo?id=2&page=1&bing=bar&bar=baz", "3"},
			},
			WhitelistQueryParams: []string{
				"id",
				"page",
			},
		},
	} {
		result, _ := FindWhitelistQueryParams(test.URLs)
		if !reflect.DeepEqual(result, test.WhitelistQueryParams) {
			t.Errorf("result %v does not match expected %v", result, test.WhitelistQueryParams)
		}
	}
}
