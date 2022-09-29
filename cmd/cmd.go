package cmd

import (
	"fmt"
	"log"
	"os"
)

// Log performs a log based on where verbosity is a selected option
func Log(v bool, s string, param ...any) {
	if v {
		log.Printf(s, param...)
	}
}

// Fatal performs a fatal log
func Fatal(s string, param ...any) {
	log.Printf(s, param...)
}

// StartUp takes the user to the homepage of the CLI tool
func StartUp() {
	if _, err := os.Stdout.WriteString(title); err != nil {
		Fatal("## " + err.Error())
	}
	fmt.Println(page, filePage, networkPage, fileForensics, fileUtility, networkForensics, networkUtility)
}
