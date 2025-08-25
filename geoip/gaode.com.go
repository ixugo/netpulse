package geoip

import (
	"context"
	"fmt"
)

// gaodeInfo 包含高德地图IP定位API返回的IP地理位置信息的完整结构体
//
//	{
//		"status": "1",
//		"info": "OK",
//		"infocode": "10000",
//		"province": "北京市",
//		"city": "北京市",
//		"adcode": "110000",
//		"rectangle": "116.0119343,39.66127144;116.7829835,40.2164962"
//	}
type gaodeInfo struct {
	Status    string `json:"status"`    // 状态值 0:失败; 1:成功
	Info      string `json:"info"`      // 状态说明 失败原因 or "OK"
	Infocode  string `json:"infocode"`  // 状态码 10000 表示正确
	Province  string `json:"province"`  // 省份 非法及国外IP地址是 无数据
	City      string `json:"city"`      // 城市 非法及国外IP地址是 无数据
	Adcode    string `json:"adcode"`    // 城市的 adcode
	Rectangle string `json:"rectangle"` // 左下右上对标对
}

func (g *gaodeInfo) toInfo() *Info {
	return &Info{
		IP:         "",   // 高德API不返回IP地址
		Country:    "中国", // 高德地图主要针对中国IP
		Region:     g.Province,
		RegionCode: "",
		City:       g.City,
		CityCode:   g.Adcode,
		ISP:        "",
		Address:    g.Province + " " + g.City,
	}
}

// Gaode 实现高德地图IP定位API
type Gaode struct {
	key string
}

// NewGaode 创建Gaode实例
func NewGaode(key string) IPer {
	return &Gaode{
		key: key,
	}
}

// Lookup 获取IP地理位置信息
func (g *Gaode) Lookup(ctx context.Context, ip string) (*Info, error) {
	// 高德地图IP定位API需要key参数，这里使用一个示例key
	// 实际使用时应该从配置中获取
	link := fmt.Sprintf("https://restapi.amap.com/v3/ip?key=%s&ip=%s", g.key, ip)
	var out gaodeInfo
	err := request(ctx, link, &out, nil)
	if err != nil {
		return nil, err
	}

	// 检查API响应状态
	if out.Status != "1" || out.Infocode != "10000" {
		return nil, fmt.Errorf("高德地图API错误: %s (状态码: %s)", out.Info, out.Infocode)
	}

	return out.toInfo(), nil
}
