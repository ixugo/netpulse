package geoip

import (
	"time"
)

var _ Cacher = (*IPCache)(nil)

type IPCache struct {
	data *TTLMap[string, *Info]
	ttl  time.Duration
}

func NewGeoIPCache(ttl time.Duration) Cacher {
	return &IPCache{
		data: NewTTLMap[string, *Info]().SetTickerCleanup(10 * time.Minute),
		ttl:  ttl,
	}
}

// Get implements Cacher.
func (g *IPCache) Get(ip string) (*Info, error) {
	info, ok := g.data.Load(ip)
	if ok {
		return info, nil
	}
	return info, ErrNotFound
}

// Set implements Cacher.
func (g *IPCache) Set(ip string, info *Info) {
	g.data.Store(ip, info, g.ttl)
}
