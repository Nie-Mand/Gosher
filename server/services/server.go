package services

import (
	"Nie-Mand/Gosher/server/core"
	"Nie-Mand/Gosher/server/schemas"
	"errors"
	"fmt"
)

type FilePings struct {
	Who         string
	FileName    string
	Description string
}

var pingForFileChannel = make(chan FilePings)
var pingForFileResponsesChannel = make(map[string]chan FilePings)

func createPingForFileChannelBridge(description string) func() {
	pingForFileResponsesChannel[description] = make(chan FilePings)

	return func() {
		delete(pingForFileResponsesChannel, description)
	}
}

func listenForPingForFileChannel(cb func(FilePings)) {
	for {
		msg := <-pingForFileChannel
		fmt.Println(msg)
		cb(msg)
	}
}

func emitToPingForFileChannel(payload FilePings) {
	pingForFileChannel <- payload
}

func listenForFilePingResponsesChannel(description string, cb func(FilePings)) {
	for {
		fmt.Println("Listening for file ping responses")
		msg := <-pingForFileResponsesChannel[description]
		fmt.Println(msg)
		cb(msg)
	}
}

func emitToFilePingResponsesChannel(description string, payload FilePings) {
	pingForFileResponsesChannel[description] <- payload
}

type ServerStruct struct {
	schemas.UnimplementedGosherServer
}

// func (s *ServerStruct) SayHi(ctx context.Context, payload *schemas.Request) (*schemas.Response, error) {
// 	who := core.GetUser(ctx)

// 	resp := schemas.Response{
// 		Msg:  payload.Msg + ", going to be sent to " + payload.Destination + " from " + who,
// 		Name: payload.Name,
// 		File: payload.File,
// 	}

// 	if channel, ok := bridge[payload.Destination]; ok {

// 		if channel == nil {
// 			resp.Msg = "User is not online, sending failed"
// 		} else {
// 			channel <- resp
// 		}

// 	} else {
// 		resp.Msg = "User is not online, sending failed"
// 	}

// 	return &resp, nil
// }

// func (s *ServerStruct) ReceiveHi(_ *schemas.Identity, server schemas.Gosher_ReceiveHiServer) error {
// 	who := core.GetUser(server.Context())
// 	fmt.Println("User " + who + " is online")

// 	bridge[who] = make(chan schemas.Response)

// 	for {
// 		select {
// 		case msg := <-bridge[who]:
// 			err := server.Send(&msg)
// 			if err != nil {
// 				return err
// 			}
// 		case closeId := <-closeBridge:
// 			if closeId == who {
// 				return nil
// 			}
// 		}
// 	}
// }

func (s *ServerStruct) PingForFile(payload *schemas.PingForFileRequest, server schemas.Gosher_PingForFileServer) error {
	who := core.GetUser(server.Context())
	description := payload.Description

	bye := createPingForFileChannelBridge(description)
	defer bye()

	fmt.Println("User " + who + " is requesting file with desciption: " + description)
	emitToPingForFileChannel(FilePings{
		Who:         who,
		FileName:    "",
		Description: description,
	})

	listenForFilePingResponsesChannel(description, func(msg FilePings) {
		fmt.Println("Sending response to " + who)
		if msg.Description == description && msg.FileName != "" {
			server.Send(&schemas.PingForFileResponse{
				FileName: msg.FileName,
				Who:      msg.Who,
			})
		}
	})

	return nil
}

func (s *ServerStruct) ListenForFilePings(server schemas.Gosher_ListenForFilePingsServer) error {
	who := core.GetUser(server.Context())
	fmt.Println("User " + who + " is listening for file pings")

	listenForPingForFileChannel(func(msg FilePings) {
		if msg.FileName == "" {
			err := server.Send(&schemas.ListenForFilePingsResponse{
				Description: msg.Description,
			})
			if err != nil {
				fmt.Println(err)
			}
			response, err := server.Recv()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Sending response to " + who)
				emitToFilePingResponsesChannel(msg.Description, FilePings{
					Who:         who,
					FileName:    response.FileName,
					Description: response.Description,
				})
			}

		}
	})
	return nil
}

type FileRequest struct {
	Who      string
	FileName string
	File     []byte
	Msg      string
}

var fileProvidersChannel = make(map[string]chan FileRequest)

func createFileProvidersChannel(who string) func() {
	fileProvidersChannel[who] = make(chan FileRequest)
	return func() {
		delete(fileProvidersChannel, who)
	}
}

func listenForFileProvidersChannel(who string, cb func(FileRequest)) {
	fmt.Println(who + ": Started Listening for file requests")
	for {
		msg := <-fileProvidersChannel[who]
		fmt.Println(msg)
		cb(msg)
	}
}

func emitToFileProvidersChannel(who string, payload FileRequest) error {
	fmt.Println("ok")
	fmt.Println(who)

	if _, ok := fileProvidersChannel[who]; ok {
		fileProvidersChannel[who] <- payload
	} else {
		return errors.New("User is not online")
	}

	return nil
}

var fileRequestChannel = make(map[string]chan FileRequest)

func createFileRequestChannel(filename string) func() {
	fileRequestChannel[filename] = make(chan FileRequest)

	return func() {
		delete(fileRequestChannel, filename)
	}
}

func listenForFileRequestChannel(filename string, cb func(FileRequest)) {
	for {
		msg := <-fileRequestChannel[filename]
		fmt.Println(msg)
		cb(msg)
	}
}

func emitToFileRequestChannel(filename string, payload FileRequest) {
	fileRequestChannel[filename] <- payload
}

// func listenForFileResponseChannel(description string, cb func(FileRequest)) {
// 	for {
// 		fmt.Println("Listening for file request responses")
// 		msg := <-fileResponseChannel[description]
// 		fmt.Println(msg)
// 		cb(msg)
// 	}
// }

// func emitToFileResponseChannel(description string, payload FileRequest) {
// 	fileResponseChannel[description] <- payload
// }

func (s *ServerStruct) RequestFile(payload *schemas.RequestFileRequest, server schemas.Gosher_RequestFileServer) error {
	who := core.GetUser(server.Context())
	filename := payload.FileName
	from := payload.Who

	bye := createFileRequestChannel(filename)
	defer bye()

	fmt.Println("User " + who + " is requesting the file " + filename + " from " + from)
	emitToFileProvidersChannel(from, FileRequest{
		Who:      from,
		FileName: filename,
		File:     nil,
	})

	listenForFileRequestChannel(filename, func(msg FileRequest) {
		fmt.Println("received file back")
		server.Send(&schemas.RequestFileResponse{
			Msg:  msg.Msg,
			File: msg.File,
		})
	})

	return nil
}

func (s *ServerStruct) SeedFile(server schemas.Gosher_SeedFileServer) error {
	who := core.GetUser(server.Context())
	fmt.Println("User " + who + " is listening for file requests")
	bye := createFileProvidersChannel(who)
	defer bye()

	listenForFileProvidersChannel(who, func(msg FileRequest) {
		fmt.Println("seed file", msg)
		server.Send(&schemas.SeedFileResponse{
			FileName: msg.FileName,
			Who:      who,
		})

		response, err := server.Recv()

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Sending response to " + who)
			if response.Msg != "" {
				emitToFileRequestChannel(msg.FileName, FileRequest{
					Who:      who,
					FileName: msg.FileName,
					Msg:      response.Msg,
					File:     nil,
				})
			} else {
				emitToFileRequestChannel(msg.FileName, FileRequest{
					Who:      who,
					FileName: msg.FileName,
					Msg:      response.Msg,
					File:     response.File,
				})
			}
		}

		// 	if err != nil {
		// 		fmt.Println(err)
		// 	}
		// 	response, err := server.Recv()
		// 	if err != nil {
		// 		fmt.Println(err)
		// 	} else {
		// 		fmt.Println("Sending response to " + who)
		// 		emitToFilePingResponsesChannel(msg.Description, FilePings{
		// 			Who:         who,
		// 			FileName:    response.FileName,
		// 			Description: response.Description,
		// 		})
		// 	}

		// }
	})

	// who := core.GetUser(server.Context())
	// fmt.Println("User " + who + " is listening for file pings")

	// for {
	// 	msg := <-pingForFileBridge
	// 	fmt.Println(msg)
	// 	if msg.FileName == "" {

	// 		err := server.Send(&schemas.ListenForFilePingsResponse{
	// 			Description: msg.Description,
	// 		})

	// 		if err != nil {
	// 			fmt.Println(err)

	// 		}
	// 		response, err := server.Recv()
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		} else {
	// 			pingForFileBridge <- FilePings{
	// 				Who:         who,
	// 				FileName:    response.FileName,
	// 				Description: response.Description,
	// 			}
	// 		}

	// 	}
	// }

	return nil
}

func StartServer() {
	port := core.GetEnv("SERVER_PORT", "50051")

	server, listen := core.CreateServer(port)

	schemas.RegisterGosherServer(server, &ServerStruct{})

	listen()
}
