FROM golang:1.20.1-alpine AS dl

WORKDIR /download

RUN apk add --no-cache unzip && \
  wget -q "https://github.com/protocolbuffers/protobuf/releases/download/v22.0/protoc-22.0-linux-x86_64.zip" -O "protobuf.zip" && \
  unzip -o protobuf.zip -d protobuf && chmod -R 755 protobuf/*

FROM golang:1.20.1-alpine AS builder

COPY --from=dl /download/protobuf/bin /usr/local/bin/
