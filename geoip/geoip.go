package geoip

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

type (
	Language        string
	WrapBodyHandler func(io.Reader) io.Reader
)

const (
	English = Language("en")
	Chinese = Language("zh-CN")
)

type IPer interface {
	Lookup(ctx context.Context, ip string) (*Info, error)
}

type Cacher interface {
	Get(string) (*Info, error)
	Set(string, *Info)
}

type Engine struct {
	language Language
	handlers []IPer
	cache    Cacher
}

func New(language Language, opts ...Option) *Engine {
	e := Engine{
		language: language,
		cache:    NewGeoIPCache(time.Hour),
	}

	switch language {
	case English:
		e.handlers = append(e.handlers, NewIPapi(), NewFreeIPAPI(), NewIfconfigco(), NewIPwho())
	case Chinese:
		e.handlers = append(e.handlers, NewWhoisPconline())
	}

	for _, opt := range opts {
		opt(&e)
	}

	return &e
}

func (e *Engine) Lookup(ctx context.Context, ip string) (info *Info, err error) {
	netip := net.ParseIP(ip)
	if netip == nil {
		return nil, errors.New("invalid ip")
	}
	if netip.IsPrivate() {
		return nil, ErrPrivateIP
	}

	if e.cache != nil {
		info, err = e.cache.Get(ip)
		if err == nil {
			return
		}
	}

	for _, handler := range e.handlers {
		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		info, err = handler.Lookup(ctx, ip)
		cancel()
		if err == nil {
			if e.cache != nil {
				e.cache.Set(ip, info)
			}
			return
		}
	}
	return info, err
}

type Info struct {
	IP         string
	Country    string // Country
	Region     string // Province/State
	RegionCode string // Province/State code
	City       string // City
	CityCode   string // City code
	ISP        string // Internet Service Provider
	Address    string // Address (e.g., "Hubei Province Jingmen City China Unicom")
}

func request(ctx context.Context, link string, out any, wrapBody WrapBodyHandler) error {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, link, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code: %d", resp.StatusCode)
	}
	if wrapBody != nil {
		return json.NewDecoder(wrapBody(resp.Body)).Decode(&out)
	}
	return json.NewDecoder(resp.Body).Decode(&out)
}
