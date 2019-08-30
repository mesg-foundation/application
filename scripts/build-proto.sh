#!/bin/bash

# make sure script is running inside mesg/tools container.
source $(dirname $0)/require-mesg-tools.sh

PROJECT=/project
GRPC=$PROJECT/protobuf
GRPC_PLUGIN="--go_out=plugins=grpc,paths=source_relative:."

protoc $GRPC_PLUGIN --proto_path=$PROJECT $GRPC/types/event.proto
protoc $GRPC_PLUGIN --proto_path=$PROJECT $GRPC/types/execution.proto
protoc $GRPC_PLUGIN --proto_path=$PROJECT $GRPC/types/instance.proto
protoc $GRPC_PLUGIN --proto_path=$PROJECT $GRPC/types/process.proto

protoc --gogofaster_out=plugins=grpc,paths=source_relative:./service/ \
       --proto_path=/project/protobuf/types \
       service.proto

protoc $GRPC_PLUGIN --proto_path=$PROJECT $GRPC/api/event.proto
protoc $GRPC_PLUGIN --proto_path=$PROJECT $GRPC/api/execution.proto
protoc $GRPC_PLUGIN --proto_path=$PROJECT $GRPC/api/instance.proto
protoc $GRPC_PLUGIN --proto_path=$PROJECT $GRPC/api/service.proto
protoc $GRPC_PLUGIN --proto_path=$PROJECT $GRPC/api/process.proto

protoc $GRPC_PLUGIN --proto_path=$PROJECT $GRPC/coreapi/api.proto
