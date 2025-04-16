#!/bin/bash
CURRENT_DIR=$(pwd)
for x in $(find ${CURRENT_DIR}/Mini-Tweeter-Protocol-Buffer/* -type d); do
  protoc -I=${x} -I=${CURRENT_DIR}/Mini-Tweeter-Protocol-Buffer -I /usr/local/include --go_out=${CURRENT_DIR} \
   --go-grpc_out=${CURRENT_DIR} ${x}/*.proto
done
