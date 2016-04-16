package main

import (
	"github.com/paked/configure"
)

var (
	config               = configure.New()
	configBind           = config.String("bind", ":5327", "Bind with IP address and port")
	configDomainName     = config.String("domain-name", "serf.", "Domain name")
	configSerfRPCAddress = config.String("serf", "127.0.0.1:7373", "Serf RPC Address")
	configSerfRPCAuthKey = config.String("serf-auth", "", "Serf RPC auth key")
)

func init() {
	config.Use(configure.NewFlagWithUsage(usage))
	config.Use(configure.NewEnvironment())
	config.Use(configure.NewFlag())
}

func usage() string {
	return `Usage: serf-dns [options]
Options:
-h               This help
--bind           Bind to interface and port (default: 0.0.0.0:5327)
--domain-name    Specify domain name (default: serf.)
--serf           Serf RPC address (default: 127.0.0.1:7373)
--serf-auth      Serf RPC authentication key (default: empty)
`
}
