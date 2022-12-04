package services

import (
	"Nie-Mand/Gosher/server/core"
	"Nie-Mand/Gosher/server/schemas"
	"context"
)

var bridge = make(map[string](chan string))
var closeBridge = make(chan string)

type ServerStruct struct {
	schemas.UnimplementedGosherServer
}

func (s *ServerStruct) SayHi(ctx context.Context, payload *schemas.Request) (*schemas.Response, error) {
	resp := schemas.Response{
		Msg: payload.Msg + ", going to be sent to " + payload.Destination,
	}

	if channel, ok := bridge[payload.Destination]; ok {

		if channel == nil {
			resp.Msg = "User is not online, sending failed"
		} else {
			channel <- resp.Msg
		}

	} else {
		resp.Msg = "User is not online, sending failed"
	}

	return &resp, nil
}

func (s *ServerStruct) ReceiveHi(payload *schemas.Identity, server schemas.Gosher_ReceiveHiServer) error {
	bridge[payload.Id] = make(chan string)

	for {
		select {
		case msg := <-bridge[payload.Id]:
			err := server.Send(&schemas.Response{Msg: msg})
			if err != nil {
				return err
			}
		case closeId := <-closeBridge:
			if closeId == payload.Id {
				return nil
			}
		}
	}
}

func StartServer() {
	port := core.GetEnv("SERVER_PORT", "50051")

	server, listen := core.CreateServer(port)

	schemas.RegisterGosherServer(server, &ServerStruct{})

	listen()
}
