package main

import (
	"net"
	"testing"
)

func TestNewHostRecord(t *testing.T) {
	IPAddress := net.ParseIP("192.3.4.5")
	host := newHostRecord("web.role.serf", IPAddress, 0)

	if host == nil {
		t.Errorf("Failed to create new host record")
	}
}
