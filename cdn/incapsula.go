package cdn

import (
	"bytes"
	"io/ioutil"
	"net"
	"net/http"
)

const incapsulaIPUrl = "https://my.incapsula.com/api/integration/v1/ips"

// LoadIncapsula loads the IP range of Incapsula CDN
func LoadIncapsula() ([]*net.IPNet, error) {
	// First get IPv4 range
	resp, err := http.Post(incapsulaIPUrl, "application/x-www-form-urlencoded", bytes.NewBuffer([]byte("resp_format=text")))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// parse and get ipv4
	return getIPRangeFromText(string(body)), nil
}
