package main

import (
	"log"
	"myz-torrent/server"
)

func main() {
	s := &server.Server{}

	log.Panic(s.Run())
}
