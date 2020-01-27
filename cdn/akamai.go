package cdn

import "net"

var akamaiASN = []string{"AS12222"}

// LoadAkamai loads the IP range of akamai CDN by looking up the ASN number
func LoadAkamai() ([]*net.IPNet, error) {
	// First get IPv4 range
	return getIPRangeFromASNNumbers(akamaiASN), nil
}
