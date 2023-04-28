package chitu

import (
	"context"
	"fmt"
	"guanyu/server"
	"log"
)

type chitu struct {
	logger *log.Logger
	server.UnimplementedGuanyuServer
}

func NewChiTu(logger *log.Logger) *chitu {
	//logger.SetPrefix("[]:")
	return &chitu{logger: logger}
}

func (c *chitu) Login(ctx context.Context, in *server.LoginRequest) (*server.LoginResponse, error) {
	c.logger.Printf("auth login --> username %s, password %s", in.Username, in.Password)
	if in.Username == "admin" && in.Password == "admin" {
		c.logger.Print("auth login success.")
		return &server.LoginResponse{AuthRes: true, Active: true, PermissionList: []*server.Permission{
			{Id: "func1", Enable: true}, {Id: "func2", Enable: true},
		}}, nil
	}
	c.logger.Printf("failed to auth login.")
	return nil, fmt.Errorf("login failed")
}
