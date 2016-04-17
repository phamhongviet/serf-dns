#!/bin/bash

SERF=quay.io/phamhongviet/serf

LEADER=`docker run -d -p 127.0.0.1:7373:7373 $SERF agent -rpc-addr '0.0.0.0:7373'`
echo $LEADER
LEADER_IP=`docker inspect -f '{{.NetworkSettings.IPAddress}}' $LEADER`

TEMP_CONFIG=`mktemp`
cat > $TEMP_CONFIG <<'EOF'
{
  "rpc_addr": "0.0.0.0:7373",
  "rpc_auth": "IX2Uzr/UQ3nrdM7U6wMBFA=="
}
EOF
docker run -d -p 127.0.0.1:7374:7373 -v ${TEMP_CONFIG}:/etc/serf.json $SERF agent -config-file /etc/serf.json -join $LEADER_IP

docker run -d $SERF agent -tag role=web -tag dc=cali -tag srv=foo -join $LEADER_IP
docker run -d $SERF agent -tag role=web -tag dc=cali -tag srv=foo -join $LEADER_IP
docker run -d $SERF agent -tag role=web -tag dc=cali -tag srv=bar -join $LEADER_IP

docker run -d $SERF agent -tag role=web -tag dc=oreg -tag srv=bar -join $LEADER_IP
docker run -d $SERF agent -tag role=web -tag dc=oreg -tag srv=foo -join $LEADER_IP

docker run -d $SERF agent -tag role=db -tag dc=cali -join $LEADER_IP
docker run -d $SERF agent -tag role=db -tag dc=oreg -join $LEADER_IP

DEAD=`docker run -d $SERF agent -tag role=web -tag dc=oreg -tag srv=bar -join $LEADER_IP`
DEAD="$DEAD `docker run -d $SERF agent -tag role=web -tag dc=cali -tag srv=foo -join $LEADER_IP`"
sleep 0.5
docker stop -t 0 $DEAD > /dev/null
docker rm $DEAD > /dev/null
