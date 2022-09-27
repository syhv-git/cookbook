package cmd

import (
	"log"
)

func Log(v bool, s string, param ...any) {
	if v {
		log.Printf(s, param...)
	}
}
