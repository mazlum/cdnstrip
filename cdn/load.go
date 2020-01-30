package cdn

import (
	"net"
	"sync"
)

// LoadAll loads and returns all CDN IP ranges
func LoadAll() ([]*net.IPNet, error) {

	var wg sync.WaitGroup
	var allRanges []*net.IPNet
	cidrChan := make(chan []*net.IPNet, 8)
	errChan := make(chan error, 8)

	wg.Add(1)
	go func() {
		cidr, err := LoadCloudFront()
		cidrChan <- cidr
		errChan <- err
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		cidr, err := LoadDDOSGuard()
		cidrChan <- cidr
		errChan <- err
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		cidr, err := LoadFastly()
		cidrChan <- cidr
		errChan <- err
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		cidr, err := LoadAzure()
		cidrChan <- cidr
		errChan <- err
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		cidr, err := LoadAkamai()
		cidrChan <- cidr
		errChan <- err
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		cidr, err := LoadCloudflare()
		cidrChan <- cidr
		errChan <- err
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		cidr, err := LoadMaxCdn()
		cidrChan <- cidr
		errChan <- err
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		cidr, err := LoadIncapsula()
		cidrChan <- cidr
		errChan <- err
		wg.Done()
	}()

	wg.Wait()
	close(cidrChan)
	close(errChan)

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
