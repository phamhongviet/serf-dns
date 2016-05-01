package main

import (
	"testing"
)

func TestCheckCustomDomainNameExistence(t *testing.T) {
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

	for dn := range SFTable {
		ok := checkCustomDomainNameExistence(dn, SFTable)
		if !ok {
			t.Errorf("Failed to find existing custom domain name %s", dn)
		}
	}

	for _, dn := range []string{"not.exist.serf.", "no.good.nope.", "also.not.exists."} {
		ok := checkCustomDomainNameExistence(dn, SFTable)
		if ok {
			t.Errorf("Failed: Custom domain name %s actually does not exist", dn)
		}
	}
}
