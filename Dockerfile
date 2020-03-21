FROM golang:1.13.8-buster AS builder
LABEL name=myz-torrent
LABEL maintainer="myz <xuziye0327@gmail.com>"

ENV SERVER_ADDR 0.0.0.0
ENV SERVER_PORT 8080
ENV DOWNLOAD_PATH "~/myz_torrent_download/"
ENV CONFIG_FILE=
ENV ARGS=

WORKDIR /go/src/myz-torrent
COPY . .
RUN go build

ENTRYPOINT ./myz-torrent \
    -s ${SERVER_ADDR} \
    -p ${SERVER_PORT} \
    -d ${DOWNLOAD_PATH} \
    ${ARGS}

EXPOSE ${SERVER_PORT}
