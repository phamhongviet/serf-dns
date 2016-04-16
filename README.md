# serf-dns

[![Build Status](https://travis-ci.org/phamhongviet/serf-dns.svg?branch=master)](https://travis-ci.org/phamhongviet/serf-dns)

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

### Test
To test the project:

```
make test
```

### Build
To build the binary executable file `serf-dns`:

```
make build
```

### Play
After building `serf-dns`, you can play with it:

```
./serf-dns
```

And in another terminal, send some DNS requests to it:

```
dig @localhost -p 5327 web.role.serf
dig @localhost -p 5327 db.role.serf
dig @localhost -p 5327 oreg.dc.web.role.serf
dig @localhost -p 5327 db.role.cali.dc.serf
dig @localhost -p 5327 foo.srv.serf
dig @localhost -p 5327 bar.srv.serf
```

And you can compare the result using `serf members`

### Clean
Testing and playing with the project need a few serf agents running in docker containers. To clean those up, run:

```
make clean
```

## Configuration

Serf-dns can be configured using environment variables or command line parameters. If both environment variable and parameter for the same directive are provided, serf-dns will use environment variable.

* __bind__: bind with IP address and port.         
Default: Bind all interfaces with port 5327.         
Use environment variable `BIND` or parameter `--bind=`.         
For example, to bind localhost with port 5300:           
```
BIND='127.0.0.1:5300' ./serf-dns
./serf-dns --bind='127.0.0.1:5300'
```

* __domain-name__: domain name for serf hosts        
Default: `serf.`       
Use environment variable `DOMAIN_NAME` or parameter `--domain-name=`.            
For example, to use domain name `my.dn.`:              
```
DOMAIN_NAME='my.dn.' ./serf-dns
./serf-dns --domain-name='my.dn.'
```
_Note_: It is important to have `.` at the end of domain name

* __serf__: serf RPC address            
Default: 127.0.0.1:7373            
Use environment variable `SERF` or parameter `--serf=`.            
For example, if serf agent is run at localhost port 8000, use:
```
SERF='127.0.0.1:8000' ./serf-dns
./serf-dns --serf='127.0.0.1:8000'
```

## TODO

* Support hostname query
* Support configuration with file
* Support authentication in serf agent
* Clean and test functions in main.go
