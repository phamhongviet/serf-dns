package main

type serfFilterTable map[string]serfFilter

func checkCustomDomainNameExistence(domainName string, sftab serfFilterTable) bool {
	_, ok := sftab[domainName]
	return ok
}

func loadCustomDomainName(data []byte) serfFilterTable {
	return serfFilterTable{}
}
