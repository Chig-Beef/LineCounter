package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

var errors int
var totalLine int
var fileTypes []string

var giveDetails bool

func main() {
	args := os.Args

	args = removeFlags(args)

	var path string
	fileTypes, path = parseArgs(args)

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

func parseArgs(args []string) ([]string, string) {
	if len(args) < 3 {
		log.Fatal(
			"line counter takes a minimum of 2 parameters.\n" +
				"lc fileTypes... path",
		)
	}
	return args[1 : len(args)-1], args[len(args)-1]
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

	if !slices.Contains(fileTypes, getFileType(file.Name())) {
		return
	}

	data, err := os.ReadFile(path)
	if err != nil {
		errors++
		return
	}

	lineCount := len(strings.Split(string(data), "\n"))

	if giveDetails {
		fmt.Println(path, lineCount)
	}

	totalLine += lineCount
}

func removeFlags(args []string) []string {
	newArgs := []string{}

	for _, arg := range args {
		if arg[0] != '-' {
			newArgs = append(newArgs, arg)
		} else {
			switch arg {
			case "-details":
				giveDetails = true
			}
		}
	}

	return newArgs
}
