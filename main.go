package main

import (
	"guanyu/server"
	"guanyu/server/chitu"
	"log"
)

func main() {
	logger := log.Default()
	ct := chitu.NewChiTu(logger)
	gy := server.NewGuanyu(logger)
	go gy.Start(ct)
	select {}
}
