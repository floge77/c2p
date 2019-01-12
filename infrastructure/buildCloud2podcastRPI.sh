#!/bin/bash

echo "_____->Building Cloud2Podcast Application<-_____"
docker build -t cloud2podcastrpi -f ../DockerfileRPI ../
