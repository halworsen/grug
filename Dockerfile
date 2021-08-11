FROM golang:1.16-buster AS build

WORKDIR /go/src/grug
COPY . .

RUN go mod download
RUN go mod tidy

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN go build -v -o ./grug ./discord

FROM alpine:3.14.0

RUN apk add --no-cache mysql-client
RUN apk add bash

RUN addgroup -g 1000 grug
RUN adduser -u 1000 -G grug -h /home/grug -s /bin/bash -D grug
WORKDIR /home/grug

# configuration and commands are not copied over and must be supplied with configmaps
COPY --from=build /go/src/grug/grug .

USER grug
ENTRYPOINT [ "./grug" ]