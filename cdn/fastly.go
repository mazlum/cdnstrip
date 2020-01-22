package cdn

import "net"

const fastlyIPUrl = "https://api.fastly.com/public-ip-list"

type fastlyResponse struct {
	Addresses     []string `json:"addresses"`
	Ipv6Addresses []string `json:"ipv6_addresses"`
}

// LoadFastly loads the IP range of fastly CDN
func LoadFastly() ([]*net.IPNet, error) {
	var fastly fastlyResponse
	response, err := getJSONFromURL(fastlyIPUrl, fastly)
	if err != nil {
		return nil, err
	}
	// get ipv4 addresses
	ranges := getIPRangeFromArray(response["addresses"])
	// get regional ip ranges and return
	return append(ranges, getIPRangeFromArray(response["ipv6_addresses"])...), nil
}
