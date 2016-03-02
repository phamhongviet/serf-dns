package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	serf_client "github.com/hashicorp/serf/client"
	"github.com/miekg/dns"
)

const defaultAddr = ":5327"
const defaultDomainName = "serf."
const defaultSerfRPCAddress = "127.0.0.1:7373"

func handle(writer dns.ResponseWriter, request *dns.Msg, serfClient serf_client.RPCClient) {
	message := new(dns.Msg)
	message.SetReply(request)
	/*
		TODO
		turn questions into serf filters
		for each serf filters, get a set of hosts from serf agent
		create and add answers from the above set of hosts
	*/
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
	// connect to serf agent
	serfClient, err := client.NewRPCClient(defaultSerfRPCAddress)
	if err != nil {
		fmt.Println(err.Error())
	}

	dns.HandleFunc(defaultDomainName,
		func(writer dns.ResponseWriter, request *dns.Msg) {
			handle(writer, request, serfClient)
		})
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
