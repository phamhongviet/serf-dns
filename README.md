# serf-dns
DNS service with [serf](https://www.serfdom.io "Hashicorp's Serf") backend

## Goal
Provide a DNS interface to query serf's host, either by name or by tags. For example:

To query redis servers in Oregon (or any host with tags service=redis and region=oregon):

```
$ dig +short redis.service.oregon.region.serf
192.168.10.11
192.168.10.12
192.168.10.13
```

Tags' order doesn't matter

```
$ dig +short oregon.region.redis.service.serf
192.168.10.11
192.168.10.13
192.168.10.12
```

Or to query a server with hostname node7.examp.le:

```
$ dig +short node7.examp.le.name.serf
192.168.7.17
```

## Develop
This project is currently develop in Golang 1.6 with `docker` and `make`

To test the project:

```
make test
```

To build the binary executable file `serf-dns`:

```
make build
```
