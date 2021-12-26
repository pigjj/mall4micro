#!/usr/bin/env bash

cd mall4micro-auth/protos && protoc --go_out=../grpc_dto --go_opt=paths=source_relative --go-grpc_out=../grpc_dto --go-grpc_opt=paths=source_relative *.proto
#cd mall4micro-auth/protos && protoc -I=. --go-grpc_out=../grpc_dto --go-grpc_opt=paths=source_relative *.proto