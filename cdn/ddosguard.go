package cdn

import "net"

const ddosGuardASN = "AS262254"

// LoadDDOSGuard loads the IP range of ddos-guard.net CDN by looking up the ASN number
func LoadDDOSGuard() ([]*net.IPNet, error) {
	// First get IPv4 range
	body, err := getTextFromURL("https://api.hackertarget.com/aslookup/?q=" + ddosGuardASN)
	if err != nil {
		return nil, err
	}
	// parse and get ipv4
	return getIPRangeFromText(body), nil
}
