package geoip

import "context"

// freeIPapiInfo contains the complete IP geolocation information struct returned by free.freeipapi.com API
// Note: Response content about China information may be inaccurate
//
//	{
//		"ipVersion": 4,
//		"ipAddress": "183.95.255.255",
//		"latitude": 39.9175,
//		"longitude": 116.362,
//		"countryName": "China",
//		"countryCode": "CN",
//		"capital": "Beijing",
//		"phoneCodes": [
//		  86
//		],
//		"timeZones": [
//		  "Asia/Shanghai",
//		  "Asia/Urumqi"
//		],
//		"zipCode": "430022",
//		"cityName": "Jinrongjie (Xicheng District)",
//		"regionName": "Beijing",
//		"continent": "Asia",
//		"continentCode": "AS",
//		"currencies": [
//		  "CNY"
//		],
//		"languages": [
//		  "zh"
//		],
//		"asn": "4837",
//		"asnOrganization": "CHINA UNICOM China169 Backbone",
//		"isProxy": false
//	  }
type freeIPapiInfo struct {
	IPVersion       int      `json:"ipVersion"`       // IP version (4 or 6)
	IPAddress       string   `json:"ipAddress"`       // IP address
	Latitude        float64  `json:"latitude"`        // Latitude
	Longitude       float64  `json:"longitude"`       // Longitude
	CountryName     string   `json:"countryName"`     // Country name
	CountryCode     string   `json:"countryCode"`     // Country code
	Capital         string   `json:"capital"`         // Capital city
	PhoneCodes      []int    `json:"phoneCodes"`      // Phone area codes
	TimeZones       []string `json:"timeZones"`       // Time zone list
	ZipCode         string   `json:"zipCode"`         // Postal code
	CityName        string   `json:"cityName"`        // City name
	RegionName      string   `json:"regionName"`      // Region name
	Continent       string   `json:"continent"`       // Continent name
	ContinentCode   string   `json:"continentCode"`   // Continent code
	Currencies      []string `json:"currencies"`      // Currency list
	Languages       []string `json:"languages"`       // Language list
	ASN             string   `json:"asn"`             // ASN number
	ASNOrganization string   `json:"asnOrganization"` // ASN organization name
	IsProxy         bool     `json:"isProxy"`         // Whether it's a proxy IP
}

func (f *freeIPapiInfo) toInfo() *Info {
	return &Info{
		IP:      f.IPAddress,
		Country: f.CountryName,
		Region:  f.RegionName,
		City:    f.CityName,
		ISP:     f.ASNOrganization,
		Address: f.RegionName + " " + f.CityName,
	}
}

type FreeIPAPI struct{}

func NewFreeIPAPI() IPer {
	return &FreeIPAPI{}
}

// Lookup implements IPer.
func (f *FreeIPAPI) Lookup(ctx context.Context, ip string) (*Info, error) {
	const link = "https://free.freeipapi.com/api/json/"
	var out freeIPapiInfo
	err := request(ctx, link, ip, &out, nil)
	return out.toInfo(), err
}
