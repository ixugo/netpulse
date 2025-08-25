package geoip

import (
	"context"
	"errors"
	"io"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// whoisPconlineInfo 包含whois.pconline.com.cn API返回的IP地理位置信息的完整结构体
//
//	{
//	    "ip": "183.95.255.255",
//	    "pro": "湖北省",
//	    "proCode": "420000",
//	    "city": "荆门市",
//	    "cityCode": "420800",
//	    "region": "",
//	    "regionCode": "0",
//	    "addr": "湖北省荆门市 联通",
//	    "regionNames": "",
//	    "err": ""
//	}
type whoisPconlineInfo struct {
	IP          string `json:"ip"`          // IP地址
	Pro         string `json:"pro"`         // 省份名称
	ProCode     string `json:"proCode"`     // 省份代码
	City        string `json:"city"`        // 城市名称
	CityCode    string `json:"cityCode"`    // 城市代码
	Region      string `json:"region"`      // 地区
	RegionCode  string `json:"regionCode"`  // 地区代码
	Addr        string `json:"addr"`        // 完整地址
	RegionNames string `json:"regionNames"` // 地区名称
	Err         string `json:"err"`         // 错误信息
}

func (i *whoisPconlineInfo) toInfo() *Info {
	return &Info{
		IP:         i.IP,
		Country:    "",
		Region:     i.Pro,
		RegionCode: i.ProCode,
		City:       i.City,
		CityCode:   i.CityCode,
		Address:    i.Addr,
	}
}

type whoisPconline struct {
	wrapBody WrapBodyHandler
}

func NewWhoisPconline() IPer {
	return &whoisPconline{
		wrapBody: func(r io.Reader) io.Reader {
			return transform.NewReader(r, simplifiedchinese.GB18030.NewDecoder())
		},
	}
}

func (w *whoisPconline) Lookup(ctx context.Context, ip string) (*Info, error) {
	const link = `http://whois.pconline.com.cn/ipJson.jsp?json=true&ip=`
	var out whoisPconlineInfo
	err := request(ctx, link, ip, &out, w.wrapBody)
	if err != nil {
		return nil, err
	}
	if out.Err != "" {
		return nil, errors.New(out.Err)
	}
	return out.toInfo(), nil
}
