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

Or to query a pre-configured domain name `us-web.serf` with name `web-.*` and tag `dc=us-.*`:

```
$ dig +short us-web.serf
192.168.6.17
192.168.7.18
192.168.8.29
192.168.9.105
```
_Note_: please see custom-domain-name.md


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

* __serf-auth__: serf RPC auth key           
Default: empty        
Use environment variable `SERF_AUTH` or parameter `--serf-auth=`        
For example, if serf agent RPC interface require key `S3creTT0k3n`, use:
```
SERF_AUTH='S3creTT0k3n' ./serf-dns
./serf-dns --serf-auth='S3creTT0k3n'
```

* __custom__: path to custom domain name file        
Default: empty        
Use environment variable `CUSTOM` or parameter `--custom=`        
For example, to load custom domain names from /etc/serf-dns/custom.json:       
```
CUSTOM='/etc/serf-dns/custom.json' ./serf-dns
./serf-dns --custom='/etc/serf-dns/custom.json'
```
Please see custom-domain-name.md for more information.

## TODO

* Support configuration with file
* Clean and test functions in main.go
* Proper logging
* Support custom TTL
