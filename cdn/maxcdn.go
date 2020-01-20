package cdn

import (
	"fmt"
	"net"
)

const maxCdnIPUrl = "https://www.maxcdn.com/one/assets/ips.txt"


// loads the IP range of Max CDN
func LoadMaxCdn() ([]*net.IPNet, error) {
	// First get IPv4 range
	fmt.Println("Getting Max CDN ip ranges")
	body, err := getIpTextFromUrl(maxCdnIPUrl)
	if err != nil {
		return nil, err
	}
	// parse and get ipv4
	return GetIPRangeFromText(body), nil
}
