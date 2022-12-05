package services

import (
	"Nie-Mand/Gosher/server/core"
	"Nie-Mand/Gosher/server/schemas"
	"context"
	"fmt"
)

var bridge = make(map[string](chan schemas.Response))

// var pingForFileBridge = make(chan string)
var closeBridge = make(chan string)

type ServerStruct struct {
	schemas.UnimplementedGosherServer
}

func (s *ServerStruct) SayHi(ctx context.Context, payload *schemas.Request) (*schemas.Response, error) {
	who := core.GetUser(ctx)

	resp := schemas.Response{
		Msg:  payload.Msg + ", going to be sent to " + payload.Destination + " from " + who,
		Name: payload.Name,
		File: payload.File,
	}

	if channel, ok := bridge[payload.Destination]; ok {

		if channel == nil {
			resp.Msg = "User is not online, sending failed"
		} else {
			channel <- resp
		}

	} else {
		resp.Msg = "User is not online, sending failed"
	}

	return &resp, nil
}

func (s *ServerStruct) ReceiveHi(_ *schemas.Identity, server schemas.Gosher_ReceiveHiServer) error {
	who := core.GetUser(server.Context())
	fmt.Println("User " + who + " is online")

	bridge[who] = make(chan schemas.Response)

	for {
		select {
		case msg := <-bridge[who]:
			err := server.Send(&msg)
			if err != nil {
				return err
			}
		case closeId := <-closeBridge:
			if closeId == who {
				return nil
			}
		}
	}
}

func (s *ServerStruct) PingForFile(payload *schemas.PingForFileRequest, server schemas.Gosher_PingForFileServer) error {
	who := core.GetUser(server.Context())
	description := payload.Description

	// pingForFileBridge[description] = make(chan string)

	resp := schemas.PingForFileResponse{
		Who:      "Pikaboo",
		Metadata: "This is a file from " + who + " with description: " + description,
	}

	server.Send(&resp)

	// if channel, ok := bridge[payload.Destination]; ok {

	// 	if channel == nil {
	// 		resp.Msg = "User is not online, sending failed"
	// 	} else {
	// 		channel <- resp.Msg
	// 	}

	// } else {
	// 	resp.Msg = "User is not online, sending failed"
	// }

	return nil
}

func (s *ServerStruct) ListenForFilePings(server schemas.Gosher_ListenForFilePingsServer) error {
	// who := core.GetUser(server.Context())

	return nil
}

func StartServer() {
	port := core.GetEnv("SERVER_PORT", "50051")

	server, listen := core.CreateServer(port)

	schemas.RegisterGosherServer(server, &ServerStruct{})

	listen()
}
