package main

import (
	"net"
	"testing"
)

func TestNewHostRecord(t *testing.T) {
	IPAddress := net.ParseIP("192.3.4.5")
	host := newHostRecord("web.role.serf", IPAddress, 7)

	if host.A.String() != "192.3.4.5" {
		t.Errorf("Failed to create new host record, wrong IP address.")
	}
	if host.Hdr.Name != "web.role.serf" {
		t.Errorf("Failed to create new host record, wrong name.")
	}
	if host.Hdr.Ttl != 7 {
		t.Errorf("Failed to create new host record, wrong TTL.")
	}
}
