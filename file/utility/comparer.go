package utility

import (
	"log"
	"reflect"
	"strings"
	"time"
)

func compareDir(x, y any) bool {
	i := reflect.ValueOf(x).String()
	j := reflect.ValueOf(y).String()
	p := strings.Split(i, "/")
	q := strings.Split(j, "/")
	if len(p) > len(q) {
		return true
	}

	n := strings.Compare(i, j)
	if n == 1 {
		return true
	}
	return false
}

func compareMod(x, y any) bool {
	i := reflect.ValueOf(x).String()
	j := reflect.ValueOf(y).String()

	a, err := time.Parse(time.RFC3339Nano, i)
	if err != nil {
		log.Fatal(err.Error())
	}
	b, err := time.Parse(time.RFC3339Nano, j)
	if err != nil {
		log.Fatal(err.Error())
	}

	if a.After(b) {
		return true
	}
	return false
}

func compareName(x, y any) bool {
	i := reflect.ValueOf(x).String()
	j := reflect.ValueOf(y).String()

	n := strings.Compare(i, j)
	if n == 1 {
		return true
	}
	return false
}

func compareSize(x, y any) bool {
	i := reflect.ValueOf(x).Int()
	j := reflect.ValueOf(y).Int()

	if i > j {
		return true
	}
	return false
}
