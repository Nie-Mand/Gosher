package services

import (
	"fmt"

	"Nie-Mand/Gosher/server/schemas"
	"Nie-Mand/Gosher/server/utils"
)

/*
* This file contains the files search services
 */

type FilePings struct {
	Who         string
	FileName    string
	Description string
}

var pingForFileListenersCount int = 0
var pingForFileChannel = make(chan FilePings)
var pingForFileResponsesChannel = make(map[string]chan FilePings)

/*
* @function: PingForFile
* @description: Search for a file
* @params: payload: *schemas.PingForFileRequest, server: schemas.Gosher_PingForFileServer
* @returns: error
 */
func (s *ServerStruct) PingForFile(payload *schemas.PingForFileRequest, server schemas.Gosher_PingForFileServer) error {
	who := utils.GetUser(server.Context())
	description := payload.Description

	bye := _createPingForFileChannelBridge(
		&pingForFileResponsesChannel,
		description,
	)

	defer bye()

	fmt.Println("User " + who + " is requesting file with desciption: " + description)
	_emitToPingForFileChannel(
		&pingForFileChannel,
		&pingForFileListenersCount,
		FilePings{
			Who:         who,
			FileName:    "",
			Description: description,
		},
	)

	_listenForFilePingResponsesChannel(
		&pingForFileResponsesChannel,
		description,
		func(msg FilePings) {
			fmt.Println("Sending response to " + who)
			if msg.Description == description && msg.FileName != "" {
				server.Send(&schemas.PingForFileResponse{
					FileName: msg.FileName,
					Who:      msg.Who,
				})
			}
		},
	)

	return nil
}

/*
* @function: ListenForFilePings
* @description: Listen for file pings
* @params: server: schemas.Gosher_ListenForFilePingsServer
* @returns: error
 */
func (s *ServerStruct) ListenForFilePings(server schemas.Gosher_ListenForFilePingsServer) error {
	who := utils.GetUser(server.Context())
	pingForFileListenersCount++
	fmt.Println("User " + who + " is listening for file pings")
	fmt.Printf("Total listeners: %d\n", pingForFileListenersCount)

	_listenForPingForFileChannel(
		&pingForFileChannel,
		func(msg FilePings) {
			if msg.FileName == "" {
				err := server.Send(&schemas.ListenForFilePingsResponse{
					Description: msg.Description,
				})
				if err != nil {
					fmt.Println(err)
				}
				for {

					response, err := server.Recv()
					if err != nil {
						fmt.Println(err)
					} else {
						fmt.Println("Sending response to " + who)
						_emitToFilePingResponsesChannel(
							&pingForFileResponsesChannel,
							msg.Description,
							FilePings{
								Who:         who,
								FileName:    response.FileName,
								Description: response.Description,
							},
						)
					}

				}
			}
		},
	)
	return nil
}

// Private functions

func _createPingForFileChannelBridge(
	pingForFileResponsesChannel *map[string]chan FilePings,
	description string,
) func() {
	(*pingForFileResponsesChannel)[description] = make(chan FilePings)

	return func() {
		delete((*pingForFileResponsesChannel), description)
	}
}

func _listenForPingForFileChannel(
	pingForFileChannel *chan FilePings,
	cb func(FilePings),
) {
	for {
		msg := <-(*pingForFileChannel)
		fmt.Println(msg)
		cb(msg)
	}
}

func _emitToPingForFileChannel(
	pingForFileChannel *chan FilePings,
	pingForFileListenersCount *int,
	payload FilePings,
) {
	for i := 0; i < (*pingForFileListenersCount); i++ {
		(*pingForFileChannel) <- payload
	}
}

func _listenForFilePingResponsesChannel(
	pingForFileResponsesChannel *map[string]chan FilePings,
	description string,
	cb func(FilePings),
) {
	for {
		fmt.Println("Listening for file ping responses")
		msg := <-(*pingForFileResponsesChannel)[description]
		fmt.Println(msg)
		cb(msg)
	}
}

func _emitToFilePingResponsesChannel(
	pingForFileResponsesChannel *map[string]chan FilePings,
	description string,
	payload FilePings,
) {
	(*pingForFileResponsesChannel)[description] <- payload
}
