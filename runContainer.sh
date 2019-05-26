#!/bin/bash
set -o nounset
set -o errexit

if [[ "$#" -ne 1 ]]; then
    echo -n "Docker Container with name cloud2podcast will be started with the `pwd`/downloads mounted"
    echo ""
    echo -n "If the reuired is not clear to you please visit https://github.com/floge77/cloud2podcast#download-tracks"
    exit 1
fi

echo -n "_____->Running cloud2podcast<-_____"
echo ""
echo -n "Checking IP of Host machine"

HOST_IP=`ifconfig | grep -Eo 'inet (addr:)?([0-9]*\.){3}[0-9]*' | grep -Eo '([0-9]*\.){3}[0-9]*' | grep -v '127.0.0.1'`
echo ""
echo -n "Ip used to serve Podcasts: $HOST_IP"
echo ""
docker run --rm --name cloud2podcast -p 8080:8080 -v `pwd`/downloads:/downloads -e HOST_IP=$HOST_IP -it cloud2podcast