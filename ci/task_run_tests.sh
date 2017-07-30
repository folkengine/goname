#!/bin/sh

set -e

# task script is in resource-tutorial/10_job_inputs/ folder
# application input is in gopath/src/github.com/cloudfoundry-community/simple-go-web-app folder
# $GOPATH is gopath/ folder
export GOPATH=$PWD/go
export PATH=$GOPATH/bin:$PATH

sudo add-apt-repository ppa:masterminds/glide && sudo apt-get update
sudo apt-get install glide

glide install

go test ./...