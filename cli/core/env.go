package core

import (
	"os"

	"github.com/joho/godotenv"
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

/*
* @function: LoadDotEnv
* @description: Load the .env file
* @params: none
* @return: none
 */
func LoadDotEnv() {
	godotenv.Load()
}
