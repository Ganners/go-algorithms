package simplest

import (
	"net/url"
)

type Response struct {
	URL      string
	Checksum uint64
}

// The goal is to find the set of query params that actually make the
// difference in the outcome of the Response
func FindWhitelistQueryParams(urls []Response) []string {
	checksumInverse := make(map[uint64][]url.Values)
	// fill the inverse with query values
	for _, response := range urls {
		parsed, err := url.Parse(response.URL)
		if err != nil {
			return nil
		}
		checksumInverse[response.Checksum] = append(checksumInverse[response.Checksum], parsed.Query())
	}

	globalWhitelist := make(map[string]struct{}, 0)
	blacklist := make(map[string]struct{}, 0)

	// loop through each checksum with multiple entries
	for _, urls := range checksumInverse {
		histogram := map[string][]string{}
		for _, requestParams := range urls {
			for key, values := range requestParams {
				value := values[0]
				if _, ok := histogram[key]; !ok {
					histogram[key] = make([]string, 0)
				}
				histogram[key] = append(histogram[key], value)
			}
		}

		// keep if all items in a histogram are equal and they match the number
		// of urls for the checksum
	filterHistogram:
		for key, values := range histogram {
			if len(values) != len(urls) {
				blacklist[key] = struct{}{}
				continue
			}
			checkValue := values[0]
			for _, value := range values {
				if value != checkValue {
					blacklist[key] = struct{}{}
					continue filterHistogram
				}
			}
		}
		for key := range histogram {
			if _, ok := blacklist[key]; !ok {
				globalWhitelist[key] = struct{}{}
			}
		}
	}

	whitelist := make([]string, 0, len(globalWhitelist))
	for key := range globalWhitelist {
		whitelist = append(whitelist, key)
	}
	return whitelist
}
