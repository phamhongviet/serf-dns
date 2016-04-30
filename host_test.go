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
			Addr: net.ParseIP("192.3.4.5"),
		},
		{
			Addr: net.ParseIP("192.4.5.6"),
		},
	}
	domainName := "web.role.serf."

	message.Answer = addHostsToAnswer(hosts, domainName, message.Answer)

	if len(message.Answer) != len(hosts) {
		t.Errorf("Failed to add hosts to answer. Want: %d. Get: %d", len(hosts), len(message.Answer))
		t.FailNow()
	}

	for i, rr := range message.Answer {
		if rr.Header().Name != domainName {
			t.Errorf("Wrong hostname is added in answer.")
		}

		rrA, ok := rr.(*dns.A)
		if ok {
			if rrA.A.String() != hosts[i].Addr.String() {
				t.Errorf("Wrong host address is added in answer")
			}
		}
	}
}
