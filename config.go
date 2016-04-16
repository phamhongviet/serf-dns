package main

import (
	"github.com/paked/configure"
)

var (
	config               = configure.New()
	configBind           = config.String("bind", ":5327", "Bind with IP address and port")
	configDomainName     = config.String("domain-name", "serf.", "Domain name")
	configSerfRPCAddress = config.String("serf", "127.0.0.1:7373", "Serf RPC Address")
)

func init() {
	config.Use(configure.NewEnvironment())
	config.Use(configure.NewFlag())
}
