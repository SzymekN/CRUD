# syntax=docker/dockerfile:1

## Build
FROM golang:alpine3.16 as build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd/*.go ./
COPY Makefile ./
COPY pkg/ ./pkg

RUN apk add --no-cache git
RUN go install github.com/go-swagger/go-swagger/cmd/swagger
RUN swagger generate spec -o /swagger.json --scan-models
# RUN make swagger
RUN go build -o /userapi

## Deploy
FROM alpine:3.16
WORKDIR /
COPY --from=build /swagger.json /docs/
COPY --from=build /userapi /userapi
EXPOSE 8200
ENTRYPOINT [ "/userapi" ]
