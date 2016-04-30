package main

import (
	"net"
	"testing"

	serf_client "github.com/hashicorp/serf/client"
	"github.com/miekg/dns"
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
	if host.Hdr.Rrtype != dns.TypeA {
		t.Errorf("Failed to create new host record, wrong RR type.")
	}
	if host.Hdr.Class != dns.ClassINET {
		t.Errorf("Failed to create new host record, wrong class.")
	}
	if host.Hdr.Ttl != 7 {
		t.Errorf("Failed to create new host record, wrong TTL.")
	}
}

func TestAddHostsToAnswer(t *testing.T) {
	message := new(dns.Msg)
	hosts := []serf_client.Member{
		{
			Name: "web.role.serf.",
			Addr: net.ParseIP("192.3.4.5"),
		},
		{
			Name: "web.role.serf.",
			Addr: net.ParseIP("192.4.5.6"),
		},
	}
	message.Answer = addHostsToAnswer(hosts, message.Answer)
}
