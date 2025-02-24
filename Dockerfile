FROM golang:1.23.3 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Change CGO_ENABLED on 0 when no sqlite3 drive
RUN CGO_ENABLED=1 GOOS=linux go build -o /placelists-back ./cmd/server


# FROM gcr.io/distroless/base-debian11 AS build-release-stage
FROM golang:1.23.3 AS build-release-stage

COPY --from=build-stage /placelists-back /placelists-back

EXPOSE 8082

CMD ["/placelists-back"]