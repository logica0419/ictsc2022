FROM builder:v1 AS builder

WORKDIR /build
COPY ./go.* ./
RUN go mod download

COPY ./protobuf ./protobuf/
RUN protoc -I . --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative protobuf/ping.proto

COPY . .
RUN go build -o /app main.go

FROM alpine:3.14.2

COPY --from=builder /app /app

ENTRYPOINT [ "/app" ]
