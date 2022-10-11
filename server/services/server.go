package services

import (
	"Nie-Mand/Gosher/server/core"
	out "Nie-Mand/Gosher/server/schemas"
	"context"
)

type ServerStruct struct {
	out.UnimplementedHelloWorldServer
}

func (s *ServerStruct) SayHi(ctx context.Context, custom *services.Custom) (*services.CustomResponse, error) {
	resp := out.CustomResponse{
		Msg: custom.Msg + ", but from server",
	}
	return &resp, nil
}

func StartServer() {
	port := core.GetEnv("SERVER_PORT", "50051")

	server, listen := core.CreateServer(port)

	out.RegisterHelloWorldServer(server, &ServerStruct{})

	listen()
}
