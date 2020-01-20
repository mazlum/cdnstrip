package cdn

import (
	"net"
)

const maxCdnIPUrl = "https://www.maxcdn.com/one/assets/ips.txt"

// LoadMaxCdn loads the IP range of Max CDN
func LoadMaxCdn() ([]*net.IPNet, error) {
	// First get IPv4 range
	body, err := getTextFromURL(maxCdnIPUrl)
	if err != nil {
		return nil, err
	}
	// parse and get ipv4
	return getIPRangeFromText(body), nil
}
