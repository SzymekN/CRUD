# syntax=docker/dockerfile:1

## Build
FROM golang:alpine3.16 as build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd/*.go ./
COPY pkg/ ./pkg

RUN go build -o /userapi

## Deploy
FROM alpine:latest
WORKDIR /
COPY --from=build /userapi /userapi
EXPOSE 8200
ENTRYPOINT [ "/userapi" ]
