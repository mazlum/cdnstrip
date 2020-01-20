package cdn

import (
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

// LoadCloudflare loads the IP range of cloudflare CDN
func LoadCloudflare() ([]*net.IPNet, error) {

	// First get IPv4 range
	resp, err := http.Get("https://www.cloudflare.com/ips-v4")
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ranges []*net.IPNet
	for _, i := range strings.Split(string(body), "\n") {
		_, cidr, err := net.ParseCIDR(i)
		if err != nil {
			return nil, err
		}
		ranges = append(ranges, cidr)
	}

	// Then get IPv6 range
	resp, err = http.Get("https://www.cloudflare.com/ips-v6")
	if err != nil {
		return nil, err
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	for _, i := range strings.Split(string(body), "\n") {
		_, cidr, err := net.ParseCIDR(i)
		if err != nil {
			return nil, err
		}
		ranges = append(ranges, cidr)
	}

	return ranges, nil

}
