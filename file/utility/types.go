package utility

import (
	. "cookbook/file"
	"log"
	"reflect"
	"strings"
	"time"
)

type sortDir func(Node) string

type sortMod func(Node) time.Time

type sortName func(Node) string

type sortSize func(Node) int64

func (s sortDir) handle(n Node) any {
	return func(n Node) string { return n.Path }(n)
}

func (s sortMod) handle(n Node) any {
	return func(n Node) string { return n.ModTime().Format(time.RFC3339Nano) }(n)
}

func (s sortName) handle(n Node) any {
	return func(n Node) string { return n.Name() }(n)
}

func (s sortSize) handle(n Node) any {
	return func(n Node) int64 { return n.Size() }(n)
}

func (s sortDir) compare(x, y any) bool {
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

func (s sortMod) compare(x, y any) bool {
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

func (s sortName) compare(x, y any) bool {
	i := reflect.ValueOf(x).String()
	j := reflect.ValueOf(y).String()

	n := strings.Compare(i, j)
	if n == 1 {
		return true
	}
	return false
}

func (s sortSize) compare(x, y any) bool {
	i := reflect.ValueOf(x).Int()
	j := reflect.ValueOf(y).Int()

	if i > j {
		return true
	}
	return false
}

type constraint interface {
	sortDir | sortMod | sortName | sortSize

	compare(x, y any) bool
	handle(n Node) any
}
