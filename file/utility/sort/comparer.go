package sort

import (
	cmd "cookbook"
	"log"
	"reflect"
	"strings"
	"time"
)

func compareDir(v bool, x, y any) bool {
	i := reflect.ValueOf(x).String()
	j := reflect.ValueOf(y).String()
	cmd.Log(v, "- Comparing dir:\n%s\n%s\n", i, j)

	p := strings.Split(i, "/")
	q := strings.Split(j, "/")
	if len(p) > len(q) {
		cmd.Log(v, "- %s has more parent directories\n", i)
		return true
	}

	n := strings.Compare(i, j)
	if n == 1 {
		cmd.Log(v, "- %s comes before %s\n", i, j)
		return true
	}
	cmd.Log(v, "- %s comes before %s\n", j, i)
	return false
}

func compareMod(v bool, x, y any) bool {
	i := reflect.ValueOf(x).String()
	j := reflect.ValueOf(y).String()
	cmd.Log(v, "- Comparing last modified times:\n%v\n%v\n", x, y)

	a, err := time.Parse(time.RFC3339Nano, i)
	if err != nil {
		log.Fatal(err.Error())
	}
	b, err := time.Parse(time.RFC3339Nano, j)
	if err != nil {
		log.Fatal(err.Error())
	}

	if a.After(b) {
		cmd.Log(v, "- %v comes before %v\n", a, b)
		return true
	}
	cmd.Log(v, "- %v comes before %v\n", b, a)
	return false
}

func compareName(v bool, x, y any) bool {
	i := reflect.ValueOf(x).String()
	j := reflect.ValueOf(y).String()
	cmd.Log(v, "- Comparing names: %s %s\n", i, j)

	n := strings.Compare(i, j)
	if n == 1 {
		cmd.Log(v, "- %s comes before %s\n", i, j)
		return true
	}
	cmd.Log(v, "- %s comes before %s\n", j, i)
	return false
}

func compareSize(v bool, x, y any) bool {
	i := reflect.ValueOf(x).Int()
	j := reflect.ValueOf(y).Int()
	cmd.Log(v, "- Comparing sizes: %d %d\n", i, j)

	if i > j {
		cmd.Log(v, "- %d comes before %d\n", i, j)
		return true
	}
	cmd.Log(v, "- %d comes before %d\n", j, i)
	return false
}
