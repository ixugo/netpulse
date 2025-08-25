package geoip

import (
	"context"
	"testing"
)

func TestLookupIPWithWhoisPconline(t *testing.T) {
	e := New(Chinese)
	info, err := e.Lookup(context.Background(), "183.95.255.255")
	if err != nil {
		t.Fatal(err)
	}
	if info.IP != "183.95.255.255" {
		t.Fatal("ip not match")
	}
	if info.Region != "湖北省" {
		t.Fatal("region not match")
	}
	if info.RegionCode != "420000" {
		t.Fatal("region code not match")
	}
	if info.City != "荆门市" {
		t.Fatal("city not match")
	}
	if info.CityCode != "420800" {
		t.Fatal("city code not match")
	}
	if info.Address != "湖北省荆门市 联通" {
		t.Fatal("address not match")
	}
}

func TestLookupIPWithIfconfigco(t *testing.T) {
	e := New(English, WithHandlers(NewIfconfigco()))
	info, err := e.Lookup(context.Background(), "183.95.255.255")
	if err != nil {
		t.Fatal(err)
	}
	if info.IP != "183.95.255.255" {
		t.Fatalf("ip not match, got: %s", info.IP)
	}
	if info.Country != "China" {
		t.Fatalf("country not match, got: %s", info.Country)
	}
	if info.Region != "Hubei" {
		t.Fatalf("region not match, got: %s", info.Region)
	}
	if info.RegionCode != "HB" {
		t.Fatalf("region code not match, got: %s", info.RegionCode)
	}
	if info.City != "Wuhan" {
		t.Fatalf("city not match, got: %s", info.City)
	}
	if info.ISP != "CHINA UNICOM China169 Backbone" {
		t.Fatalf("ISP not match, got: %s", info.ISP)
	}
	expectedAddr := "China Hubei Wuhan CHINA UNICOM China169 Backbone"
	if info.Address != expectedAddr {
		t.Fatalf("address not match, got: %s, expected: %s", info.Address, expectedAddr)
	}
}

func TestLookupIPWithFreeIPAPI(t *testing.T) {
	e := New(English, WithHandlers(NewFreeIPAPI()))
	info, err := e.Lookup(context.Background(), "183.95.255.255")
	if err != nil {
		t.Fatal(err)
	}
	if info.IP != "183.95.255.255" {
		t.Fatalf("ip not match, got: %s", info.IP)
	}
	// Note: freeipapi may return inaccurate information for Chinese IPs
	if info.Country != "China" {
		t.Fatalf("country not match, got: %s", info.Country)
	}
	if info.Region != "Beijing" {
		t.Fatalf("region not match, got: %s", info.Region)
	}
	if info.City != "Jinrongjie (Xicheng District)" {
		t.Fatalf("city not match, got: %s", info.City)
	}
	if info.ISP != "CHINA UNICOM China169 Backbone" {
		t.Fatalf("ISP not match, got: %s", info.ISP)
	}
	expectedAddr := "Beijing Jinrongjie (Xicheng District)"
	if info.Address != expectedAddr {
		t.Fatalf("address not match, got: %s, expected: %s", info.Address, expectedAddr)
	}
}

func TestLookupIPWithIPwho(t *testing.T) {
	e := New(English, WithHandlers(NewIPwho()))
	info, err := e.Lookup(context.Background(), "183.95.255.255")
	if err != nil {
		t.Fatal(err)
	}
	if info.IP != "183.95.255.255" {
		t.Fatalf("ip not match, got: %s", info.IP)
	}
	if info.Country != "China" {
		t.Fatalf("country not match, got: %s", info.Country)
	}
	if info.Region != "Hubei" {
		t.Fatalf("region not match, got: %s", info.Region)
	}
	if info.RegionCode != "42" {
		t.Fatalf("region code not match, got: %s", info.RegionCode)
	}
	if info.City != "Wuhan" {
		t.Fatalf("city not match, got: %s", info.City)
	}
	if info.CityCode != "230060" {
		t.Fatalf("city code not match, got: %s", info.CityCode)
	}
	if info.ISP != "China Unicom China1 Backbone" {
		t.Fatalf("ISP not match, got: %s", info.ISP)
	}
	expectedAddr := "China Hubei Wuhan China Unicom Hubei Province Network"
	if info.Address != expectedAddr {
		t.Fatalf("address not match, got: %s, expected: %s", info.Address, expectedAddr)
	}
}

func TestLookupIPWithIPapi(t *testing.T) {
	e := New(English, WithHandlers(NewIPapi()))
	info, err := e.Lookup(context.Background(), "31.223.184.41")
	if err != nil {
		t.Fatal(err)
	}
	if info.IP != "31.223.184.41" {
		t.Fatalf("ip not match, got: %s", info.IP)
	}
	if info.Country != "Japan" {
		t.Fatalf("country not match, got: %s", info.Country)
	}
	if info.Region != "Tokyo" {
		t.Fatalf("region not match, got: %s", info.Region)
	}
	if info.RegionCode != "13" {
		t.Fatalf("region code not match, got: %s", info.RegionCode)
	}
	if info.City != "Koto-ku" {
		t.Fatalf("city not match, got: %s", info.City)
	}
	if info.ISP != "G-Core Labs" {
		t.Fatalf("ISP not match, got: %s", info.ISP)
	}
	expectedAddr := "Japan Tokyo Koto-ku G-Core Labs S.A"
	if info.Address != expectedAddr {
		t.Fatalf("address not match, got: %s, expected: %s", info.Address, expectedAddr)
	}
}
