# syntax=docker/dockerfile:1

## Build
FROM golang:alpine3.16 as build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY config/ ./config
COPY controller/ ./controller
COPY model/ ./model
COPY storage/ ./storage

RUN go build -o /userapi

## Deploy
FROM alpine:latest
WORKDIR /
COPY --from=build /userapi /userapi
EXPOSE 8200
ENTRYPOINT [ "/userapi" ]
