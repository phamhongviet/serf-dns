package main

import (
	"testing"
)

func TestParseDomainName(t *testing.T) {
	config.Parse()

	SFTable := serfFilterTable{
		"dead.digit.serf.": serfFilter{
			Name:   "^[0-9].*",
			Status: "failed",
		},
		"digit.name.serf.": serfFilter{
			Name:   "^[0-9].*",
			Status: "alive",
		},
		"dead.serf.": serfFilter{
			Status: "failed",
		},
	}

	for dn, sf := range SFTable {
		resultSF := parseDomainName(dn, SFTable)
		ok := sf.Compare(resultSF)
		if !ok {
			t.Errorf("Failed to parse custom domain name %s", dn)
		}
	}
}

func TestParseTagsDomainName(t *testing.T) {
	config.Parse()

	domainNameSample := "foo.srv.cali.dc.serf."
	expectedSerfFilter := serfFilter{
		Tags: map[string]string{
			"srv": "foo",
			"dc":  "cali",
		},
		Status: "alive",
	}
	resultSerfFilter := parseTagsDomainName(domainNameSample)
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

func TestParseCustomDomainName(t *testing.T) {
	SFTable := serfFilterTable{
		"dead.digit.serf.": serfFilter{
			Name:   "^[0-9].*",
			Status: "failed",
		},
		"digit.name.serf.": serfFilter{
			Name:   "^[0-9].*",
			Status: "alive",
		},
		"dead.serf.": serfFilter{
			Status: "failed",
		},
	}

	for dn, sf := range SFTable {
		resultSF := parseCustomDomainName(dn, SFTable)
		ok := sf.Compare(resultSF)
		if !ok {
			t.Errorf("Failed to parse custom domain name %s", dn)
		}
	}
}
