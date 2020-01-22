package cdn

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

//
func getIPRangeFromArray(arr []string) []*net.IPNet {
	var ranges []*net.IPNet
	// split and parse cidr addresses
	for _, i := range arr {
		_, cidr, err := net.ParseCIDR(i)
		if err == nil {
			ranges = append(ranges, cidr)
		}
	}
	return ranges
}

// getIPRangeFromText parse ip range body
func getIPRangeFromText(body string) []*net.IPNet {
	return getIPRangeFromArray(strings.Split(body, "\n"))
}

// getTextFromUrl response body with a basic GET request
func getTextFromURL(addr string) (string, error) {
	resp, err := http.Get(addr)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// get json data and Unmarshal from url
// data is a struct like cloudFrontResponse
func getJSONFromURL(addr string, data interface{}) (map[string][]string, error) {
	res, err := http.Get(addr)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &data)

	returnData := make(map[string][]string)
	for key, values := range data.(map[string]interface{}) {
		for _, value := range values.([]interface{}) {
			returnData[key] = append(returnData[key], value.(string))
		}
	}
	return returnData, err
}
