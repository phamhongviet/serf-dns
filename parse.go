package main

import (
	"strings"
)

func parseDomainName(domainName string) serfFilter {
	domainName = strings.TrimSuffix(domainName, configDomainName)

	tags := make(map[string]string)

	for domainName != "" {
		tagValue, tagName, remain := findTag(domainName)

		tags[tagName] = tagValue
		domainName = remain
	}

	sf := serfFilter{
		Name:   "",
		Status: "alive",
		Tags:   tags,
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
