FROM golang:1.12-alpine AS builder
RUN apk add --no-cache curl git openssh \
 && adduser -D -g '' appuser

COPY . /go/src/go.roomook.backend
WORKDIR /go/src/go.roomook.backend


ENV GO111MODULE=on
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

RUN  go mod tidy \
     && go mod verify \
     && go mod vendor



RUN go build -ldflags="-w -s" -o /go/src/go.roomook.backend/api

FROM alpine:latest

RUN apk add --no-cache tzdata

WORKDIR /app

ENV ENVIRONMENT="prod"

RUN apk add --no-cache ca-certificates

COPY --from=builder /go/src/go.roomook.backend/config ./config
COPY --from=builder /go/src/go.roomook.backend/api .


CMD [ "/app/api" ]

