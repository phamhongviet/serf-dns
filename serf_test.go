package main

import (
	"testing"
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
	client, err := connectSerfAgent("127.0.0.1:55555")
	defer func() {
		if client != nil {
			client.Close()
		}
	}()

	if err == nil {
		t.Errorf("Connect to wrong address return no error")
	}

	if client != nil {
		t.Errorf("Connect to wrong address return opened client")
	}
}

func TestConnectSerfAgentExpectingSuccess(t *testing.T) {
	client, err := connectSerfAgent(defaultSerfRPCAddress)
	defer func() {
		if client != nil {
			client.Close()
		}
	}()

	if err != nil {
		t.Errorf("Connect to default address return no error. Did you setup test environment?")
	}

	if client == nil {
		t.Errorf("Connect to default address return opened client. Did you setup test environment?")
	}
}