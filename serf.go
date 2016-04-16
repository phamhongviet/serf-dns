package main

import (
	"reflect"
	"strings"

	serf_client "github.com/hashicorp/serf/client"
)

type serfFilter struct {
	Tags   map[string]string
	Status string
	Name   string
}

func (sf1 *serfFilter) Compare(sf2 serfFilter) bool {
	if strings.Compare(sf1.Status, sf2.Status) != 0 {
		return false
	}

	if strings.Compare(sf1.Name, sf2.Name) != 0 {
		return false
	}

	return reflect.DeepEqual(sf1.Tags, sf2.Tags)
}

func connectSerfAgent(serfRPCAddress string) (*serf_client.RPCClient, error) {
	return serf_client.NewRPCClient(serfRPCAddress)
}

func connectSerfAgentWithAuthKey(serfRPCAddress string, serfRPCAuthKey string) (*serf_client.RPCClient, error) {
	return nil, nil
}

func closeSerfConnection(client *serf_client.RPCClient) {
	if client != nil {
		client.Close()
	}
}

func getSerfMembers(client *serf_client.RPCClient, filter serfFilter) ([]serf_client.Member, error) {
	return client.MembersFiltered(filter.Tags, filter.Status, filter.Name)
}
