#!/bin/bash
set -o nounset
set -o errexit

if [[ "$#" -ne 1 ]]; then
    echo -n "Docker Container with name cloud2podcast will be started with the `pwd`/downloads mounted"
    echo -n "If the reuired is not clear to you please visit https://github.com/floge77/cloud2podcast#download-tracks"
    exit 1
fi

echo "_____->Running cloud2podcast<-_____"

docker run --rm --name cloud2podcast -p 8080:8080 -v `pwd`/downloads:/downloads -it cloud2podcast