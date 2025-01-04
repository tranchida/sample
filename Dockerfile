#build stage
FROM docker.io/golang:alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN go get ./...
RUN go build -o /go/bin/app ./...

#final stage
FROM docker.io/busybox:latest
COPY --from=builder /go/bin/app /app
ENTRYPOINT ["/app"]
EXPOSE 8080
