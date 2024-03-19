package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var errors int
var totalLine int
var fileType string

func main() {
	args := os.Args
	if len(args) != 3 {
		log.Panicln("error: too many or too little sys args")
	}

	fileType = args[1]
	path := args[2]

	files, err := os.ReadDir(path)
	if err != nil {
		log.Panicln("could not read dir")
	}

	errors = 0
	totalLine = 0

	for _, file := range files {
		count(file, path+"/"+file.Name())
	}

	fmt.Println("Total:", totalLine, "with", errors, "errors.")
}

func getFileType(fileName string) string {
	dot := strings.Index(fileName, ".")
	if dot == -1 {
		return ""
	}
	return fileName[dot+1:]
}

func count(file os.DirEntry, path string) {
	if file.IsDir() {
		files, err := os.ReadDir(path)
		if err != nil {
			errors++
			return
		}
		for _, file := range files {
			count(file, path+"/"+file.Name())
		}
		return
	}

	if getFileType(file.Name()) != fileType {
		return
	}

	data, err := os.ReadFile(path)
	if err != nil {
		errors++
		return
	}
	totalLine += len(strings.Split(string(data), "\n"))
}
