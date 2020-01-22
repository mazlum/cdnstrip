package cdn

import (
	"net"
)

const cloudFrontIPUrl = "http://d7uri8nf7uskq.cloudfront.net/tools/list-cloudfront-ips"

type cloudFrontResponse struct {
	CLOUDFRONTGLOBALIPLIST       []string `json:"CLOUDFRONT_GLOBAL_IP_LIST"`
	CLOUDFRONTREGIONALEDGEIPLIST []string `json:"CLOUDFRONT_REGIONAL_EDGE_IP_LIST"`
}


func LoadCloudFront() ([]*net.IPNet, error)  {
		var cf cloudFrontResponse
		response, err := getJsonFromURL(cloudFrontIPUrl, cf)
		if err != nil {
			return nil, err
		}
		// get global ip ranges
		ranges := getIpRangeFromArray(response["CLOUDFRONT_GLOBAL_IP_LIST"])
		// get regional ip ranges and return
		return append(ranges, getIpRangeFromArray(response["CLOUDFRONT_REGIONAL_EDGE_IP_LIST"])...) , nil
}
