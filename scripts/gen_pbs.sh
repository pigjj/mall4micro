#!/usr/bin/env bash

protoc --go_out=./mall4micro-common/grpc_dto --go_opt=paths=source_relative --go-grpc_out=./mall4micro-common/grpc_dto --go-grpc_opt=paths=source_relative ./mall4micro-common/protos/*.proto

protoc --go_out=./mall4micro-auth/grpc_dto --go_opt=paths=source_relative --go-grpc_out=./mall4micro-auth/grpc_dto --go-grpc_opt=paths=source_relative ./mall4micro-auth/protos/*.proto

protoc --go_out=./mall4micro-user/grpc_dto --go_opt=paths=source_relative --go-grpc_out=./mall4micro-user/grpc_dto --go-grpc_opt=paths=source_relative ./mall4micro-user/protos/*.proto