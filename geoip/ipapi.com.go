package geoip

import (
	"context"
	"errors"
)

// ipapiInfo contains the complete IP geolocation information struct returned by ipapi.com API
//
//	{
//		"status": "success",
//		"country": "Japan",
//		"countryCode": "JP",
//		"region": "13",
//		"regionName": "Tokyo",
//		"city": "Koto-ku",
//		"zip": "135-0061",
//		"lat": 35.6522,
//		"lon": 139.791,
//		"timezone": "Asia/Tokyo",
//		"isp": "G-Core Labs",
//		"org": "G-Core Labs S.A",
//		"as": "AS199524 G-Core Labs S.A.",
//		"query": "31.223.184.41"
//	  }
type ipapiInfo struct {
	Status      string  `json:"status"`      // Query status, usually "success"
	Country     string  `json:"country"`     // Country name
	CountryCode string  `json:"countryCode"` // Country code
	Region      string  `json:"region"`      // Region code
	RegionName  string  `json:"regionName"`  // Region name
	City        string  `json:"city"`        // City name
	Zip         string  `json:"zip"`         // Postal code
	Lat         float64 `json:"lat"`         // Latitude
	Lon         float64 `json:"lon"`         // Longitude
	Timezone    string  `json:"timezone"`    // Time zone
	ISP         string  `json:"isp"`         // Internet Service Provider
	Org         string  `json:"org"`         // Organization name
	AS          string  `json:"as"`          // ASN and organization info
	Query       string  `json:"query"`       // Queried IP address
}

func (i *ipapiInfo) toInfo() *Info {
	return &Info{
		IP:         i.Query,
		Country:    i.Country,
		Region:     i.RegionName,
		RegionCode: i.Region,
		City:       i.City,
		CityCode:   "",
		ISP:        i.ISP,
		Address:    i.Country + " " + i.RegionName + " " + i.City + " " + i.Org,
	}
}

// IPapi implements IPer interface
type IPapi struct{}

// NewIPapi creates IPapi instance
func NewIPapi() IPer {
	return &IPapi{}
}

// Lookup retrieves IP geolocation information
func (i *IPapi) Lookup(ctx context.Context, ip string) (*Info, error) {
	const link = "http://ip-api.com/json/"
	var out ipapiInfo
	err := request(ctx, link+ip, &out, nil)
	if err != nil {
		return nil, err
	}
	if out.Status != "success" {
		return nil, errors.New("API request failed with status: " + out.Status)
	}
	return out.toInfo(), nil
}
