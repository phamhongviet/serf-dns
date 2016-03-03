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

func handle(writer dns.ResponseWriter, request *dns.Msg, serfClient *serf_client.RPCClient) {
	message := new(dns.Msg)
	message.SetReply(request)

	for _, question := range request.Question {
		filter := parseDomainName(question.Name)
		hosts, err := serfClient.MembersFiltered(filter.Tags, filter.Status, filter.Name)
		if err != nil {
			fmt.Println(err.Error())
		}
		for _, host := range hosts {
			fmt.Printf("%s -> %s\n", question.Name, host.Addr.String())
		}
	}

	/*
		TODO: create and add answers from the above set of hosts
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
	serfClient, err := connectSerfAgent(defaultSerfRPCAddress)
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
