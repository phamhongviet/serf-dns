package main

import (
	"encoding/json"
)

type serfFilterTable map[string]serfFilter

func checkCustomDomainNameExistence(domainName string, sftab serfFilterTable) bool {
	_, ok := sftab[domainName]
	return ok
}

func loadCustomDomainName(data []byte) serfFilterTable {
	var sftab serfFilterTable

	// TODO: handle error here
	json.Unmarshal(data, &sftab)
	return sftab
}
