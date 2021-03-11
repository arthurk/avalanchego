# syntax=docker/dockerfile:experimental

FROM golang:1.15.5-buster

RUN mkdir -p /go/src/github.com/corpetty

WORKDIR $GOPATH/src/github.com/corpetty/
COPY . avalanchego

WORKDIR $GOPATH/src/github.com/corpetty/avalanchego
RUN ./scripts/build.sh

RUN ln -sv $GOPATH/src/github.com/corpetty/avalanchego/ /avalanchego
