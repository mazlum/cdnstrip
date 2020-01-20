package cdn

import "net"

// LoadAll loads and returns all CDN IP ranges
func LoadAll() ([]*net.IPNet, error) {

	var allRanges []*net.IPNet
	cidrChan := make(chan []*net.IPNet, 3)
	errChan := make(chan error, 3)
	close(cidrChan)
	close(errChan)

	go func() {
		cloudflare, err := LoadCloudflare()
		cidrChan <- cloudflare
		errChan <- err
	}()

	go func() {
		maxcdn, err := LoadMaxCdn()
		cidrChan <- maxcdn
		errChan <- err
	}()

	go func() {
		incapsula, err := LoadIncapsula()
		cidrChan <- incapsula
		errChan <- err
	}()

	for err := range errChan {
		if err != nil {
			return nil, err
		}
	}

	for cidr := range cidrChan {
		allRanges = append(allRanges, cidr...)
	}

	return allRanges, nil
}
