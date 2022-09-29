package cmd

import (
	"log"
)

func Log(v bool, s string, param ...any) {
	if v {
		log.Printf(s, param...)
	}
}

func Fatal(s string, param ...any) {
	log.Printf(s, param...)
}

func StartUp() {

}
