FROM golang:1.23.3 AS build-stage

WORKDIR /app

COPY . .

RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -o ./build cmd/server/main.go

ENTRYPOINT ["build/main"]