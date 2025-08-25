package ip

import (
	"testing"
)

func TestExternalIP(t *testing.T) {
	ip, err := ExternalIP()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ip)
}

func TestInternalIP(t *testing.T) {
	ip := InternalIP()
	t.Log(ip)

	ip = localIP()
	t.Log(ip)
}
