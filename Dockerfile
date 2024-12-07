FROM golang:1.23.3 AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o goapp

FROM alpine:latest

WORKDIR /app

COPY --from=builder /build/goapp .

CMD ["/app/goapp"]
