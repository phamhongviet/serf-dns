#!/bin/bash

SERF=quay.io/phamhongviet/serf

LEADER=`docker run -d -p 127.0.0.1:7373:7373 $SERF agent -rpc-addr '0.0.0.0:7373'`
LEADER_IP=`docker inspect -f '{{.NetworkSettings.IPAddress}}' $LEADER`

docker run -d $SERF agent -tag role=web -tag dc=cali -tag srv=foo -join $LEADER_IP
docker run -d $SERF agent -tag role=web -tag dc=cali -tag srv=foo -join $LEADER_IP
docker run -d $SERF agent -tag role=web -tag dc=cali -tag srv=bar -join $LEADER_IP

docker run -d $SERF agent -tag role=web -tag dc=oreg -tag srv=bar -join $LEADER_IP
docker run -d $SERF agent -tag role=web -tag dc=oreg -tag srv=foo -join $LEADER_IP

docker run -d $SERF agent -tag role=db -tag dc=cali -join $LEADER_IP
docker run -d $SERF agent -tag role=db -tag dc=oreg -join $LEADER_IP
