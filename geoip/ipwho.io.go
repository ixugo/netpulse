package geoip

import (
	"context"
	"errors"
)

// ipwhoInfo contains the complete IP geolocation information struct returned by ipwho.io API
//
//	{
//	    "ip": "183.95.255.255",
//	    "success": true,
//	    "type": "IPv4",
//	    "continent": "Asia",
//	    "continent_code": "AS",
//	    "country": "China",
//	    "country_code": "CN",
//	    "region": "Hubei",
//	    "region_code": "42",
//	    "city": "Wuhan",
//	    "latitude": 30.593099,
//	    "longitude": 114.305393,
//	    "is_eu": false,
//	    "postal": "230060",
//	    "calling_code": "86",
//	    "capital": "Beijing",
//	    "borders": "AF,BT,HK,IN,KG,KP,KZ,LA,MM,MN,MO,NP,PK,RU,TJ,VN",
//	    "flag": {
//	        "img": "https://cdn.ipwhois.io/flags/cn.svg",
//	        "emoji": "🇨🇳",
//	        "emoji_unicode": "U+1F1E8 U+1F1F3"
//	    },
//	    "connection": {
//	        "asn": 4837,
//	        "org": "China Unicom Hubei Province Network",
//	        "isp": "China Unicom China1 Backbone",
//	        "domain": "chinaunicom.cn"
//	    },
//	    "timezone": {
//	        "id": "Asia/Shanghai",
//	        "abbr": "CST",
//	        "is_dst": false,
//	        "offset": 28800,
//	        "utc": "+08:00",
//	        "current_time": "2025-08-25T14:54:35+08:00"
//	    }
//	}
type ipwhoInfo struct {
	IP            string  `json:"ip"`             // IP address
	Success       bool    `json:"success"`        // Whether query is successful
	Type          string  `json:"type"`           // IP type
	Continent     string  `json:"continent"`      // Continent name
	ContinentCode string  `json:"continent_code"` // Continent code
	Country       string  `json:"country"`        // Country name
	CountryCode   string  `json:"country_code"`   // Country code
	Region        string  `json:"region"`         // Region name
	RegionCode    string  `json:"region_code"`    // Region code
	City          string  `json:"city"`           // City name
	Latitude      float64 `json:"latitude"`       // Latitude
	Longitude     float64 `json:"longitude"`      // Longitude
	IsEU          bool    `json:"is_eu"`          // Whether it's EU country
	Postal        string  `json:"postal"`         // Postal code
	CallingCode   string  `json:"calling_code"`   // Calling code
	Capital       string  `json:"capital"`        // Capital city
	Borders       string  `json:"borders"`        // Border countries code
	// Flag          ipwhoFlag       `json:"flag"`           // Flag information
	Connection ipwhoConnection `json:"connection"` // Connection information
	// Timezone      ipwhoTimezone   `json:"timezone"`       // Timezone information
}

// ipwhoFlag flag information struct
// type ipwhoFlag struct {
// 	Img          string `json:"img"`           // Flag image URL
// 	Emoji        string `json:"emoji"`         // Flag emoji
// 	EmojiUnicode string `json:"emoji_unicode"` // Flag Unicode encoding
// }

// ipwhoConnection connection information struct
type ipwhoConnection struct {
	ASN    int    `json:"asn"`    // ASN number
	Org    string `json:"org"`    // Organization name
	ISP    string `json:"isp"`    // ISP name
	Domain string `json:"domain"` // Domain name
}

// ipwhoTimezone timezone information struct
// type ipwhoTimezone struct {
// 	ID          string `json:"id"`           // Timezone ID
// 	Abbr        string `json:"abbr"`         // Timezone abbreviation
// 	IsDST       bool   `json:"is_dst"`       // Whether it's DST
// 	Offset      int    `json:"offset"`       // Timezone offset in seconds
// 	UTC         string `json:"utc"`          // UTC offset
// 	CurrentTime string `json:"current_time"` // Current time
// }

func (i *ipwhoInfo) toInfo() *Info {
	return &Info{
		IP:         i.IP,
		Country:    i.Country,
		Region:     i.Region,
		RegionCode: i.RegionCode,
		City:       i.City,
		CityCode:   i.Postal, // Use postal code as city code
		ISP:        i.Connection.ISP,
		Address:    i.Country + " " + i.Region + " " + i.City + " " + i.Connection.Org,
	}
}

// IPwho 实现IPer接口的结构体
type IPwho struct{}

// NewIPwho 创建IPwho实例
func NewIPwho() IPer {
	return &IPwho{}
}

// Lookup 获取IP地理位置信息
func (i *IPwho) Lookup(ctx context.Context, ip string) (*Info, error) {
	const link = "http://ipwho.is/"
	var out ipwhoInfo
	err := request(ctx, link+ip, &out, nil)
	if err != nil {
		return nil, err
	}
	if !out.Success {
		return nil, errors.New("API request failed: success is false")
	}
	return out.toInfo(), nil
}
