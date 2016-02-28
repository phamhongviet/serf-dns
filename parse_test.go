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

func TestFindTag(t *testing.T) {
	sample := "a.b.c.d."

	expectedTagValue := "a"
	expectedTagName := "b"
	expectedRemain := "c.d."

	resultTagValue, resultTagName, resultRemain := findTag(sample)

	if expectedTagName != resultTagName {
		t.Errorf("Failed to find tag in domain name")
	}
	if expectedTagValue != resultTagValue {
		t.Errorf("Failed to find tag in domain name")
	}
	if expectedRemain != resultRemain {
		t.Errorf("Failed to find tag in domain name")
	}
}
