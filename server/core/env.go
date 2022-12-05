package core

import (
	"os"
)

/*
* @function: GetEnv
* @description: Get the environment variable
* @params: envVarName: string, defaultValue: string
* @return: string
 */
func GetEnv(envVarName string, defaultValue string) string {
	port := os.Getenv(envVarName)
	if port == "" {
		port = defaultValue
	}
	return ":" + port
}
