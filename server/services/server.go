package services

import (
	"Nie-Mand/Gosher/server/core"
	"Nie-Mand/Gosher/server/schemas"
	"context"
)

type ServerStruct struct {
	schemas.UnimplementedGosherServer
}

func (s *ServerStruct) SayHi(ctx context.Context, custom *schemas.Request) (*schemas.Response, error) {
	resp := schemas.Response{
		Msg: custom.Msg + ", but from server",
	}
	return &resp, nil
}

func StartServer() {
	port := core.GetEnv("SERVER_PORT", "50051")

	server, listen := core.CreateServer(port)

	schemas.RegisterGosherServer(server, &ServerStruct{})

	listen()
}
