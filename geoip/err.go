package geoip

import "errors"

var (
	ErrPrivateIP = errors.New("private ip")
	ErrNotFound  = errors.New("not found")
)

func IsErrPrivateIP(err error) bool {
	return errors.Is(err, ErrPrivateIP)
}
