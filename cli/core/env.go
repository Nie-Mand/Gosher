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
	value := os.Getenv(envVarName)
	if value == "" {
		value = defaultValue
	}
	return value
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
