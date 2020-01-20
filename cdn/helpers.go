package cdn

import (
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

// getIPRangeFromText parse ip range body
func getIPRangeFromText(body string) []*net.IPNet {
	var ranges []*net.IPNet
	// split and parse cidr addresses
	for _, i := range strings.Split(body, "\n") {
		_, cidr, err := net.ParseCIDR(i)
		if err == nil {
			ranges = append(ranges, cidr)
		}
	}
	return ranges
}

// getTextFromUrl response body with a basic GET request
func getTextFromURL(addr string) (string, error) {
	resp, err := http.Get(addr)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
