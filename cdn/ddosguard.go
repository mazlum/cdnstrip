package cdn

import "net"

var ddosGuardASN = []string{"AS262254"}

// LoadDDOSGuard loads the IP range of ddos-guard.net CDN by looking up the ASN number
func LoadDDOSGuard() ([]*net.IPNet, error) {
	// First get IPv4 range
	return getIPRangeFromASNNumbers(ddosGuardASN), nil
}
