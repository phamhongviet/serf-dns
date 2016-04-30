package main

import (
	"net"

	serf_client "github.com/hashicorp/serf/client"
	"github.com/miekg/dns"
)

func newHostRecord(name string, IP net.IP, TTL uint32) dns.A {
	host := dns.A{
		Hdr: dns.RR_Header{
			Name:     name,
			Rrtype:   dns.TypeA,
			Class:    dns.ClassINET,
			Ttl:      TTL,
			Rdlength: 0,
		},
		A: IP,
	}

	return host
}

func addHostsToAnswer(hosts []serf_client.Member, messageAnswer []dns.RR) []dns.RR {
	return nil
}
