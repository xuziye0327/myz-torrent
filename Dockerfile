
#build stage
FROM golang:1.13.8-buster AS builder
WORKDIR /go/src/myz-torrent
COPY . .
RUN go build

ENTRYPOINT ./myz-torrent
LABEL Name=myz-torrent Version=0.0.1
EXPOSE 8080