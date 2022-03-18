## Build

FROM golang:1.17-alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GO111MODULE=on \
    GOARCH=amd64 \
    GOPATH=/go

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go get -u github.com/swaggo/swag/cmd/swag
RUN swag init -g cmd/server/server.go

RUN go build

## Deploy

FROM alpine

WORKDIR /app

COPY --from=builder /app ./
COPY --from=builder /app ./

RUN chmod +x /app/go-rest-clean-boilerplate
RUN apk update && apk add tzdata

ENTRYPOINT ["/app/go-rest-clean-boilerplate"]
