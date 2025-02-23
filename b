package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func printTextFile(content string) {
	for _, char := range content {
		z01.PrintRune(char)
	}
}

func printError(fileName string) {
	errorMessage := "ERROR: open " + fileName + ": no such file or directory\n"
	for _, char := range errorMessage {
		z01.PrintRune(char)
	}
}

func main() {
	if len(os.Args) > 1 {
		for _, fileName := range os.Args[1:] {
			content, err := ioutil.ReadFile(fileName)
			if err != nil {
				printError(fileName)
				os.Exit(1)
			}
			printTextFile(string(content))
		}
	} else {
		content, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return
		}
		printTextFile(string(content))
	}
}
