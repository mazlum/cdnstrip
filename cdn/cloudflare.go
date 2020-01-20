package cdn

import (
	"fmt"
	"net"
)

const (
	cloudFlareIPv4Url string = "https://www.cloudflare.com/ips-v4"
	cloudFlareIPv6Url string = "https://www.cloudflare.com/ips-v6"
)

// LoadCloudflare loads the IP range of cloudflare CDN
func LoadCloudflare() ([]*net.IPNet, error) {

	fmt.Println("Getting cloudflare ipv4 ranges")
	// First get IPv4 range
	body, err := getIpTextFromUrl(cloudFlareIPv4Url)

	if err != nil {
		return nil, err
	}
	// parse and get ipv4
	ranges := GetIPRangeFromText(body)

	fmt.Println("Getting cloudflare ipv6 ranges")
	// Then get IPv6 range
	body, err = getIpTextFromUrl(cloudFlareIPv6Url)
	if err != nil {
		return nil, err
	}
	// get and append ipv6 ranges
	ranges = append(ranges, GetIPRangeFromText(body)...)

	return ranges, nil
}


