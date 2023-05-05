package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

/*
* @function: GetEnv
* @description: Get the environment variable
* @params: envVarName: string, defaultValue: string
* @return: string
 */
func GetEnv(envVarName string, defaultValue string) string {
	value := os.Getenv(envVarName)
	if value == "" {
		value = defaultValue
	}
	return value
}