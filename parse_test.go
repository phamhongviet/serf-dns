package main

import (
	"testing"
)

func TestParseDomainName(t *testing.T) {
	domainNameSample := "foo.srv.cali.dc.serf."
	expectedSerfFilter := serfFilter{
		Tags: map[string]string{
			"srv": "foo",
			"dc":  "cali",
		},
		Status: "alive",
	}
	resultSerfFilter := parseDomainName(domainNameSample)
	ok := expectedSerfFilter.Compare(resultSerfFilter)
	if !ok {
		t.Errorf("Failed to parse domain name %s", domainNameSample)
	}
}
