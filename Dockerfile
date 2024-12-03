FROM golang:1.22-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

# Copy go files
COPY go.mod go.sum main.go ./

RUN go mod download
RUN go mod tidy
RUN go build -o main .

FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /build/main .

# Copy static files
COPY index.html .
COPY script.js .

ENTRYPOINT ["/app/main"]
