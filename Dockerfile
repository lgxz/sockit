FROM golang:alpine AS build

WORKDIR /sockit
COPY main.go /sockit/

RUN apk add --update git
RUN go get golang.org/x/net/proxy && go build -ldflags="-s -w"

FROM alpine:latest
MAINTAINER "Kevin Leo <lgx@rorz.eu.org>"

COPY --from=build /sockit/sockit /usr/local/bin/

ARG port
ENV SOCKS5_PROXY="tor:9050" PORT=${port:-5327} 

EXPOSE ${PORT}
COPY entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
