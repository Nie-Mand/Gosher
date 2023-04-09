package utils

import (
	"context"

	"google.golang.org/grpc/metadata"
)

/*
* @function: GetUser
* @description: Get the user from the context
* @params: ctx: context.Context
* @return: string
 */
 func GetUser(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	if len(md["who"]) == 0 {
		return ""
	}

	return md["who"][0]
}