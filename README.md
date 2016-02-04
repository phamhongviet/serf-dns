# serf-dns
DNS service with [serf](https://www.serfdom.io "Hashicorp's Serf") backend

## Goal
Provide a DNS interface to query serf's host, either by name or by tags. For example:

To query redis servers in Oregon (or any host with tags redis and oregon):

```
$ dig +short redis.oregon.serf
192.168.10.11
192.168.10.12
192.168.10.13
```

Tags' order doesn't matter

```
$ dig +short oregon.redis.serf
192.168.10.11
192.168.10.13
192.168.10.12
```

Or to query a server with hostname node7:

```
$ dig +short node7.serf
192.168.7.17
```
