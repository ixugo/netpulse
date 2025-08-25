package geoip

import (
	"context"
)

// ifconfigcoInfo contains the complete IP geolocation information struct returned by ifconfig.co API
//
//	{
//	    "ip": "183.95.255.255",
//	    "ip_decimal": 3076521983,
//	    "country": "China",
//	    "country_iso": "CN",
//	    "country_eu": false,
//	    "region_name": "Hubei",
//	    "region_code": "HB",
//	    "city": "Wuhan",
//	    "latitude": 30.589,
//	    "longitude": 114.2681,
//	    "time_zone": "Asia/Shanghai",
//	    "asn": "AS4837",
//	    "asn_org": "CHINA UNICOM China169 Backbone",
//	    "user_agent": {
//	        "product": "Mozilla",
//	        "version": "5.0",
//	        "comment": "(Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/139.0.0.0 Safari/537.36",
//	        "raw_value": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/139.0.0.0 Safari/537.36"
//	    }
//	}
type ifconfigcoInfo struct {
	IP         string  `json:"ip"`          // IP address
	IPDecimal  int64   `json:"ip_decimal"`  // IP address in decimal format
	Country    string  `json:"country"`     // Country name
	CountryISO string  `json:"country_iso"` // Country ISO code
	CountryEU  bool    `json:"country_eu"`  // Whether the country is in EU
	RegionName string  `json:"region_name"` // Region name
	RegionCode string  `json:"region_code"` // Region code
	City       string  `json:"city"`        // City name
	Latitude   float64 `json:"latitude"`    // Latitude
	Longitude  float64 `json:"longitude"`   // Longitude
	TimeZone   string  `json:"time_zone"`   // Time zone
	ASN        string  `json:"asn"`         // ASN number
	ASNOrg     string  `json:"asn_org"`     // ASN organization name
	// UserAgent  ifconfigcoUserAgent `json:"user_agent"`  // User agent information
}

// ifconfigcoUserAgent user agent information struct
// type ifconfigcoUserAgent struct {
// 	Product  string `json:"product"`   // Product name
// 	Version  string `json:"version"`   // Version
// 	Comment  string `json:"comment"`   // Comment
// 	RawValue string `json:"raw_value"` // Raw value
// }

func (i *ifconfigcoInfo) toInfo() *Info {
	return &Info{
		IP:         i.IP,
		Country:    i.Country,
		Region:     i.RegionName,
		RegionCode: i.RegionCode,
		City:       i.City,
		CityCode:   "", // ifconfig.co API does not provide city code
		ISP:        i.ASNOrg,
		Address:    i.Country + " " + i.RegionName + " " + i.City + " " + i.ASNOrg,
	}
}

// Ifconfigco implements IPer interface
type Ifconfigco struct{}

// NewIfconfigco creates Ifconfigco instance
func NewIfconfigco() IPer {
	return &Ifconfigco{}
}

// Lookup retrieves IP geolocation information
func (i *Ifconfigco) Lookup(ctx context.Context, ip string) (*Info, error) {
	const link = "https://ifconfig.co/json?ip="
	var out ifconfigcoInfo
	err := request(ctx, link+ip, &out, nil)
	if err != nil {
		return nil, err
	}
	return out.toInfo(), nil
}
