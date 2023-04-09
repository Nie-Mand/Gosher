package services

import (
	"Nie-Mand/Gosher/server/schemas"
	"Nie-Mand/Gosher/server/utils"
	"errors"
	"fmt"
)

/*
* This file contains the file transfering services
 */

type FileRequest struct {
	Who      string
	FileName string
	File     []byte
	Msg      string
}

var fileProvidersChannel = make(map[string]chan FileRequest)
var fileRequestChannel = make(map[string]chan FileRequest)

/*
* @function: RequestFile
* @description: Request a file from a user
* @params: payload: *schemas.RequestFileRequest, server: schemas.Gosher_RequestFileServer
* @returns: error
 */
func (s *ServerStruct) RequestFile(payload *schemas.RequestFileRequest, server schemas.Gosher_RequestFileServer) error {
	who := utils.GetUser(server.Context())
	filename := payload.FileName
	from := payload.Who

	bye := _createFileRequestChannel(filename)
	defer bye()

	fmt.Println("User " + who + " is requesting the file " + filename + " from " + from)
	_emitToFileProvidersChannel(from, FileRequest{
		Who:      from,
		FileName: filename,
		File:     nil,
	})

	_listenForFileRequestChannel(filename, func(msg FileRequest) {
		fmt.Println("received file back")
		server.Send(&schemas.RequestFileResponse{
			Msg:  msg.Msg,
			File: msg.File,
		})
	})

	return nil
}

/*
* @function: SeedFile
* @description: Listen for file requests
* @params: server: schemas.Gosher_SeedFileServer
* @returns: error
 */
func (s *ServerStruct) SeedFile(server schemas.Gosher_SeedFileServer) error {
	who := utils.GetUser(server.Context())
	fmt.Println("User " + who + " is listening for file requests")
	bye := _createFileProvidersChannel(who)
	defer bye()

	_listenForFileProvidersChannel(who, func(msg FileRequest) {
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
				_emitToFileRequestChannel(msg.FileName, FileRequest{
					Who:      who,
					FileName: msg.FileName,
					Msg:      response.Msg,
					File:     nil,
				})
			} else {
				_emitToFileRequestChannel(msg.FileName, FileRequest{
					Who:      who,
					FileName: msg.FileName,
					Msg:      response.Msg,
					File:     response.File,
				})
			}
		}

	})

	return nil
}


// Private functions

func _createFileRequestChannel(filename string) func() {
	fileRequestChannel[filename] = make(chan FileRequest)

	return func() {
		delete(fileRequestChannel, filename)
	}
}

func _emitToFileProvidersChannel(who string, payload FileRequest) error {
	fmt.Println("ok")
	fmt.Println(who)

	if _, ok := fileProvidersChannel[who]; ok {
		fileProvidersChannel[who] <- payload
	} else {
		return errors.New("User is not online")
	}

	return nil
}

func _listenForFileRequestChannel(filename string, cb func(FileRequest)) {
	for {
		msg := <-fileRequestChannel[filename]
		cb(msg)
	}
}

func _createFileProvidersChannel(who string) func() {
	fileProvidersChannel[who] = make(chan FileRequest)
	return func() {
		delete(fileProvidersChannel, who)
	}
}

func _listenForFileProvidersChannel(who string, cb func(FileRequest)) {
	fmt.Println(who + ": Started Listening for file requests")
	for {
		msg := <-fileProvidersChannel[who]
		fmt.Println(msg)
		cb(msg)
	}
}

func _emitToFileRequestChannel(filename string, payload FileRequest) {
	fileRequestChannel[filename] <- payload
}
