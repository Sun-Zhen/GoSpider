#!/bin/sh

BIN_PATH=$(cd $(dirname $0) && pwd)
HOME_PATH=$(cd ${BIN_PATH} && cd .. && pwd)
export GOPATH=${HOME_PATH}

go get gopkg.in/yaml.v2
go get github.com/golang/glog
go get github.com/PuerkitoBio/goquery
go install -v spider
${BIN_PATH}/spider -log_dir ./logs

