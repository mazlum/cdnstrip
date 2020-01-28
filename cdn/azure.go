package cdn

import "net"

var azureASN = []string{"AS53587"}

// LoadDDOSGuard loads the IP range of ddos-guard.net CDN by looking up the ASN number
func LoadAzure() ([]*net.IPNet, error) {
	// First get IPv4 range
	return getIPRangeFromASNNumbers(azureASN), nil
}


