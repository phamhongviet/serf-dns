package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/miekg/dns"
)

const defaultAddr = ":5327"
const defaultDomainName = "serf."

func handle(writer dns.ResponseWriter, request *dns.Msg) {
	message := new(dns.Msg)
	message.SetReply(request)
	rr := new(dns.A)
	rr.Hdr = dns.RR_Header{Name: "localhost.serf.", Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 0}
	rr.A = net.ParseIP("127.0.0.1").To4()
	message.Answer = append(message.Answer, rr)
	err := writer.WriteMsg(message)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func serve(net string, address string) {
	server := &dns.Server{Addr: address, Net: net, TsigSecret: nil}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Failed to setup the %s server at %s: %s\n", net, address, err.Error())
	}
}

func main() {
	dns.HandleFunc(defaultDomainName, handle)
	go serve("tcp", defaultAddr)
	go serve("udp", defaultAddr)
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
forever:
	for {
		select {
		case <-sig:
			break forever
		}
	}
}
