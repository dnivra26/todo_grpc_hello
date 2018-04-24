FROM golang:alpine AS builder
RUN apk add --no-cache --update git
ENV GOPATH=/go

COPY . $GOPATH/src/todo_grpc_hello
WORKDIR $GOPATH/src/todo_grpc_hello

RUN go get
RUN go build -v -o /go/bin/todo_grpc_hello server.go

FROM alpine:latest

COPY --from=builder /go/bin/todo_grpc_hello /usr/local/bin/todo_grpc_hello

CMD todo_grpc_hello