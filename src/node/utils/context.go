package utils

import (
	"context"

	"google.golang.org/grpc/metadata"
)

/*
* @function: GetContext
* @description: Get a context
* @params: void
* @return: context.Context
 */
 func GetContext() context.Context {
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(
		ctx,
		metadata.Pairs("who", GetEnv("WHO", "anonymous")),
	)
	return ctx
}