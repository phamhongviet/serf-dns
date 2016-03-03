package main

import (
	"net"

	"github.com/miekg/dns"
)

func newHostRecord(name string, IP net.IP, TTL uint32) dns.A {
	host := dns.A{
		Hdr: dns.RR_Header{
			Name:     "",
			Rrtype:   0,
			Class:    0,
			Ttl:      0,
			Rdlength: 0,
		},
		A: net.ParseIP("0.0.0.0"),
	}

	return host
}
