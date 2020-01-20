package cdn

import (
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

// Parse ip range body
func GetIPRangeFromText(body string) []*net.IPNet {
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


// get ip ranges from txt file url
func getIpTextFromUrl(urlAddress string) (string, error) {
	resp, err := http.Get(urlAddress)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}