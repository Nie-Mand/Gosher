package utils

import (
	"log"
	"os"
	"strings"
)

/*
* @function: GetFile
* @description: Open a file and return the file
* @params: filePath: string
* @return: []byte
 */
 func GetFile(filePath string) []byte {
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Println("Error opening file")
		panic(err)
	}
	return file
}

/*
* @function: SaveFile
* @description: Save a file into the downloads directory
* @params: filePath: string, file: []byte
* @return: void
 */
func SaveFile(filePath string, file []byte) {
	err := os.WriteFile(filePath,
		file,
		0644,
	)

	if err != nil {
		log.Println("Error saving file")
		panic(err)
	} else {
		log.Println("Saved file")
	}
}

/*
* @function: CheckFilesForMatch
* @description: Check folder for files
* @params: description: string
* @return: string[]
 */
func CheckFilesForMatch(description string) []string {
	files, err := os.ReadDir("data")
	if err != nil {
		log.Println("Error reading directory")
		panic(err)
	}

	var fileNames []string
	for _, file := range files {
		if !file.IsDir() && strings.Contains(file.Name(), description) {
			fileNames = append(fileNames, file.Name())
		}
	}

	return fileNames
}

/*
* @function: OpenFile
* @description: Open a file and return the file as a byte array
* @params: filename: string
* @return: []byte
 */
func OpenFile(filename string) ([]byte, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Println("Error opening file")
		log.Println(err)
		return nil, err
	}

	return file, nil
}