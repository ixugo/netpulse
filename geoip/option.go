package geoip

import "slices"

type Option func(*Engine)

// WithHandlers set handlers
// NewFreeIPAPI() NewIfconfigco() NewIPapi() NewIPwho()
func WithHandlers(iper ...IPer) Option {
	return func(e *Engine) {
		e.handlers = slices.Clone(iper)
	}
}

func WithCache(cache Cacher) Option {
	return func(e *Engine) {
		e.cache = cache
	}
}
