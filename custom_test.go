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

func TestLoadCustomDomainName(t *testing.T) {
	data := `{
	"my-custom-dn-1.serf": {
		"name": "^web-[0-5][0-9]",
		"status": "alive",
		"tags": {
			"role": "web"
		}
	},
	"failed.web.serf": {
		"name": "^web-.*",
		"status": "failed",
		"tags": {
			"role": "web"
		}
	},
	"us.dc.serf": {
		"tags": {
			"dc": "us-.*"
		}
	}
}`
	expect := serfFilterTable{
		"my-custom-dn-1.serf": serfFilter{
			Name:   "^web-[0-5][0-9]",
			Status: "alive",
			Tags: map[string]string{
				"role": "web",
			},
		},
		"failed.web.serf": serfFilter{
			Name:   "^web-.*",
			Status: "failed",
			Tags: map[string]string{
				"role": "web",
			},
		},
		"us.dc.serf": serfFilter{
			Tags: map[string]string{
				"dc": "us-.*",
			},
		},
	}
	result := loadCustomDomainName(data)

	if len(expect) != len(result) {
		t.Errorf("Failed to load custom domain name: result serf filter table is different from the expected one.")
	}

	for dn, sf := range expect {
		resultSF := result[dn]
		if sf.Compare(resultSF) != true {
			t.Errorf("Failed to load custom domain name: result serf filter table is different from the expected one.")
		}
	}
}
