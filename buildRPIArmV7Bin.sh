#!/usr/bin/env bash
go get -d -v && CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -a -installsuffix cgo -ldflags='-w -s'  -o cloud2podcast