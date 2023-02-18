#!/bin/sh
CURRENT_DIR=$(cd $(dirname $0); pwd)
cd $CURRENT_DIR

#protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative helloworld/helloworld.proto

protoc --go_out=./../pb *.proto --go-grpc_out=require_unimplemented_servers=false:./../pbgrpc *.proto