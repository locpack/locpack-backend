FROM golang:1.23.3-alpine AS build-stage

WORKDIR /goman/

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV GOBIN=/goman/
ENV GOOS=linux
ENV GOARCH=amd64

RUN go install -v -a -tags netgo -ldflags='-w' ./cmd/...

FROM alpine:3.17 AS release-stage

COPY --from=build-stage /goman/ .