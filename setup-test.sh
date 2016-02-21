#!/bin/bash

# serf docker image
SERF=serf

LEAD=`docker run -d $SERF agent -tag dc=cali -tag srv=foo`
LEAD_IP=`docker inspect -f '{{.NetworkSettings.IPAddress}}' $LEAD`
docker run -d $SERF agent -tag dc=cali -tag srv=foo -join $LEAD_IP
docker run -d $SERF agent -tag dc=cali -tag srv=bar -join $LEAD_IP
docker run -d $SERF agent -tag dc=oreg -tag srv=bar -join $LEAD_IP
docker run -d $SERF agent -tag dc=oreg -tag srv=foo -join $LEAD_IP
docker run -d -p 127.0.0.1:7373:7373 $SERF agent -join $LEAD_IP -rpc-addr '0.0.0.0:7373'
