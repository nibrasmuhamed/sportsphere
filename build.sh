#!/bin/bash


buildTime=$(date --utc +%FT%T.%3NZ)
cId=$(git rev-parse --short=7 HEAD)
version=$(git describe --tags --abbrev=0)-${cId}


Version="github.com/mohdjishin/sportsphere/internal/meta.Version=${version}"
BuildTime="github.com/mohdjishin/sportsphere/internal/meta.BuildTime=${buildTime}"
CommitID="github.com/mohdjishin/sportsphere/internal/meta.CommitID=${cId}"

Ldflags="-X '${Version}' -X '${BuildTime}' -X '${CommitID}'  -s -w"
Command='env CGO_ENABLED=0  GOOS=linux GOARCH=amd64 go build -v -ldflags='

echo 'Starting binary build'
echo ${Command}"${Ldflags}"
${Command}"${Ldflags}"
echo 'sportsphere build completed.'
