package main

import (
	"strings"
)

func parseDomainName(domainName string) serfFilter {
	domainName = strings.TrimSuffix(domainName, defaultDomainName)
	sf := serfFilter{
		Name:   "",
		Status: "alive",
		Tags:   map[string]string{},
	}
	return sf
}

func findTag(domainName string) (tagValue, tagName, remain string) {
	if domainName == "" {
		return "", "", ""
	}

	res := strings.SplitN(domainName, ".", 3)
	return res[0], res[1], res[2]
}
