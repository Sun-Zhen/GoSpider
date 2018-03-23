#!/bin/sh

BIN_PATH=$(cd $(dirname $0) && pwd)
HOME_PATH=$(cd ${BIN_PATH} && cd .. && pwd)
export GOPATH=${HOME_PATH}
go install -v spider
${BIN_PATH}/spider -log_dir ./logs

