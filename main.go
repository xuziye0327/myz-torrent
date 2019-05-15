package main

import (
	"fmt"
	"myz-torrent/server"
)

func main() {
	s := &server.Server{}

	if err := s.Run(); err != nil {
		fmt.Println(err)
	}
}
