package main

import (
	"fmt"
	"io/ioutil"
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
		filter := parseDomainName(question.Name, SerfFilterTable)
		hosts, err := getSerfMembers(serfClient, filter)
		if err != nil {
			fmt.Println(err.Error())
		}
		message.Answer = addHostsToAnswer(hosts, question.Name, message.Answer)
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

	if *configCustomDomainNameFile != "" {
		customDNData, err := ioutil.ReadFile(*configCustomDomainNameFile)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			SerfFilterTable = loadCustomDomainName(customDNData)
		}
	}

	serfClient, err := connectSerfAgent(*configSerfRPCAddress, *configSerfRPCAuthKey)
	defer closeSerfConnection(serfClient)
	if err != nil {
		fmt.Println(err.Error())
	}

	dns.HandleFunc(*configDomainName,
		func(writer dns.ResponseWriter, request *dns.Msg) {
			handle(writer, request, serfClient)
		})
	go serve("tcp", *configBind)
	go serve("udp", *configBind)
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
