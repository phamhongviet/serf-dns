package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	serf_client "github.com/hashicorp/serf/client"
	"github.com/miekg/dns"
)

func handle(writer dns.ResponseWriter, request *dns.Msg, serfClient *serf_client.RPCClient) {
	message := new(dns.Msg)
	message.SetReply(request)

	for _, question := range request.Question {
		filter := parseDomainName(question.Name)
		hosts, err := getSerfMembers(serfClient, filter)
		if err != nil {
			fmt.Println(err.Error())
		}
		for _, host := range hosts {
			var newRR dns.RR
			newHost := newHostRecord(question.Name, host.Addr, 0)
			newRR = &newHost
			message.Answer = append(message.Answer, newRR)
		}
	}

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
	config.Parse()

	serfClient, err := connectSerfAgent(configSerfRPCAddress)
	defer closeSerfConnection(serfClient)
	if err != nil {
		fmt.Println(err.Error())
	}

	dns.HandleFunc(configDomainName,
		func(writer dns.ResponseWriter, request *dns.Msg) {
			handle(writer, request, serfClient)
		})
	go serve("tcp", configBind)
	go serve("udp", configBind)
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
