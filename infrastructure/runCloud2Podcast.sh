#!/bin/bash
set -o nounset
set -o errexit

if [[ "$#" -ne 1 ]]; then
    echo "Usage: $(basename "$0") downloadDirectory"
    echo "Docker Container with name cloud2podcast will be started with the downloadDir mounted"
    exit 1
fi

echo "_____->Running Cloud2Podcast Application<-_____"
downloadDir=$1
docker run --rm --name c2p -p 8080:8080 -v $downloadDir:/downloads -d cloud2podcast