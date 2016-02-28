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

func findTag(domainName string) (tagName, tagValue, remain string) {
	return "", "", ""
}
