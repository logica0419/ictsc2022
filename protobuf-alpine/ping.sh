#!/bin/sh

grpcurl -plaintext :8080 app.protobuf.PingService/Ping
