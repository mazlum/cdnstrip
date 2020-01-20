package cdn

import "net"

// Check returns true if the CIDR ranges contains the given IP
func Check(cidrs []*net.IPNet, ip net.IP) bool {
	for _, cidr := range cidrs {
		if cidr.Contains(ip) {
			return true
		}
	}
	return false
}
