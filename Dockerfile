FROM golang:1.24-alpine3.21 AS build-env

ENV GOBIN=/bin/apps/

COPY go.mod go.sum .
RUN go mod download

COPY . . 
RUN go install -v ./cmd/...

FROM alpine:3.21
COPY --from=build-env /bin/apps/ /apps

WORKDIR /apps/
ENV PATH="/apps:${PATH}"
