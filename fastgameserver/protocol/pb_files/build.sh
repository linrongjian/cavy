#!/bin/sh
CURRENT_DIR=$(cd $(dirname $0); pwd)
cd $CURRENT_DIR

protoc --go_out=./../pb *.proto