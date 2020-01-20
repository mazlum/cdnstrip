package cdn

import (
	"net"
)

const (
	cloudFlareIPv4Url string = "https://www.cloudflare.com/ips-v4"
	cloudFlareIPv6Url string = "https://www.cloudflare.com/ips-v6"
)

// LoadCloudflare loads the IP range of cloudflare CDN
func LoadCloudflare() ([]*net.IPNet, error) {

	// First get IPv4 range
	body, err := getTextFromURL(cloudFlareIPv4Url)

	if err != nil {
		return nil, err
	}
	// parse and get ipv4
	ranges := GetIPRangeFromText(body)

	// Then get IPv6 range
	body, err = getTextFromURL(cloudFlareIPv6Url)
	if err != nil {
		return nil, err
	}
	// get and append ipv6 ranges
	ranges = append(ranges, GetIPRangeFromText(body)...)

	return ranges, nil
}
