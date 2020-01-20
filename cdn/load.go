package cdn

import "net"

// LoadAll loads and returns all CDN IP ranges
func LoadAll() ([]*net.IPNet, error) {

	var allRanges []*net.IPNet

	r, err := LoadCloudflare()
	if err != nil {
		return nil, err
	}

	allRanges = append(allRanges, r...)

	return allRanges, nil
}
