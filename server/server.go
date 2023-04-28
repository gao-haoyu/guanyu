package server

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type Guanyu struct {
	logger *log.Logger
}

func NewGuanyu(logger *log.Logger) *Guanyu {
	//logger.SetPrefix("[]:")
	return &Guanyu{logger: logger}
}

func (gy *Guanyu) Start(server_case GuanyuServer) {
	gy.logger.Println("start guanyu server")
	guanyuServer := grpc.NewServer()
	RegisterGuanyuServer(guanyuServer, server_case)
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		gy.logger.Printf("start guanyu listen failed : %s", err)
		return
	}
	err = guanyuServer.Serve(listener)
	if err != nil {
		gy.logger.Printf("start guanyu server failed : %s", err)
		return
	}
	gy.logger.Println("start guanyu server success.")
}
