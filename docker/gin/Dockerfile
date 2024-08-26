FROM golang:1.22-alpine

ENV WORKDIR /go/src
WORKDIR $WORKDIR

COPY api/go.mod go.mod
COPY api/go.sum go.sum

RUN go install github.com/air-verse/air@latest
