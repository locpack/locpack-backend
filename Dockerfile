FROM golang:1.23.3 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /cmd/server


FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /cmd/server /cmd/server

EXPOSE 8082

USER admin:admin

ENTRYPOINT ["/cmd/server"]