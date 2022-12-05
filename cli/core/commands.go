package core

import (
	"fmt"
	"os"

	"google.golang.org/grpc"
)

func SayHi(
	client *grpc.ClientConn,
) {

}

func Init() {
	if err := os.Mkdir("cache", os.ModePerm); err != nil {
		fmt.Println("Error creating cache directory")
		panic(err)
	} else {
		fmt.Println("Created cache directory")
	}

	if err := os.Mkdir("shared", os.ModePerm); err != nil {
		fmt.Println("Error creating shared directory")
		panic(err)

	} else {
		fmt.Println("Created shared directory")
	}

	// handle TLS Stuffs
}
