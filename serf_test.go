package main

import (
	"testing"
)

const (
	testSerfRPCAddress         = "serf:7373"
	testSerfRPCAddressWithAuth = "serf-auth:7373"
	serfAuthKey                = "IX2Uzr/UQ3nrdM7U6wMBFA=="
)

func TestSerfFilterCompare(t *testing.T) {
	sf1 := serfFilter{
		Name:   "",
		Status: "alive",
		Tags: map[string]string{
			"a": "65",
			"b": "66",
			"c": "67",
		},
	}
	sf2 := serfFilter{
		Name:   "",
		Status: "alive",
		Tags: map[string]string{
			"a": "65",
			"b": "66",
			"c": "67",
		},
	}
	sf3 := serfFilter{
		Name:   "",
		Status: "alive",
		Tags: map[string]string{
			"a": "65",
			"b": "66",
		},
	}
	if sf1.Compare(sf2) != true {
		t.Errorf("SerfFilter Compare return unexpected result.")
	}
	if sf1.Compare(sf3) != false {
		t.Errorf("SerfFilter Compare return unexpected result.")
	}
}

func TestConnectSerfAgentExpectingFailure(t *testing.T) {
	client, err := connectSerfAgent("127.0.0.1:55555", "")
	defer closeSerfConnection(client)

	if err == nil {
		t.Errorf("Connect to wrong address return no error")
	}

	if client != nil {
		t.Errorf("Connect to wrong address return opened client")
	}
}

func TestConnectSerfAgentExpectingSuccess(t *testing.T) {
	client, err := connectSerfAgent(testSerfRPCAddress, "")
	defer closeSerfConnection(client)

	if err != nil {
		t.Errorf("Connect to default address return no error. Did you setup test environment?")
	}

	if client == nil {
		t.Errorf("Connect to default address return opened client. Did you setup test environment?")
	}
}

func TestConnectSerfAgentWithAuthKey(t *testing.T) {
	client, err := connectSerfAgent(testSerfRPCAddressWithAuth, serfAuthKey)
	defer closeSerfConnection(client)

	if err != nil {
		t.Errorf("Connect to serf-auth address return error. Did you setup test environment?")
	}

	if client == nil {
		t.Errorf("Connect to serf-auth address return nil client. Did you setup test environment?")
	}

	_, err = client.Members()
	if err != nil {
		t.Errorf("Cannot retrieve member list. Wrong RPC auth key?")
	}
}

func TestGetSerfMembers(t *testing.T) {
	client, err := connectSerfAgent(testSerfRPCAddress, "")
	defer closeSerfConnection(client)
	if err != nil {
		t.Errorf("Connect to default address return no error. Did you setup test environment?")
	}

	filter := serfFilter{
		Name:   "",
		Status: "alive",
		Tags: map[string]string{
			"role": "web",
			"dc":   "cali",
		},
	}

	members, err := getSerfMembers(client, filter)
	if err != nil {
		t.Errorf("Connect to default address return no error. Did you setup test environment?")
	}
	if members == nil {
		t.Errorf("No members found from serf. Did you setup test environment?")
	}
}
