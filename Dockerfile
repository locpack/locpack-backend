FROM golang:1.23.3-bullseye AS build-stage

WORKDIR /app

COPY . .

RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -o . cmd/server/main.go

FROM ubuntu AS release-stage

COPY --from=build-stage /app/main .

ENTRYPOINT ["./main"]
