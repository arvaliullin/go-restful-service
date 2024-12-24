#!/bin/bash
# git update-index --chmod=+x ./scripts/proto-gen.sh
PROTO_DIR="api/proto"
GO_OUT_DIR="."
protoc -I=${PROTO_DIR} --go_out=${GO_OUT_DIR} --go-grpc_out=${GO_OUT_DIR} ${PROTO_DIR}/glossary.proto
