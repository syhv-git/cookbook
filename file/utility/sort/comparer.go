package sort

import (
	"github.com/syhv-git/cookbook/cmd"
	"reflect"
	"strings"
	"time"
)

/*
All compare functions evaluate x > y in its own type context and returns the expression result

@devs-if you decide to add on to this project, handleDesc is the only function you need to call on as long as you've implemented the custom sort type
*/
func handleDesc[K any, T constraint](v, b bool, x, y K, data T) bool {
	if b {
		return compare(v, x, y, data)
	}
	return compare(v, y, x, data)
}

func compare[K any, T constraint](v bool, x, y K, data T) bool {
	t := reflect.ValueOf(data)
	d := reflect.Indirect(t)
	i := d.Interface()
	switch i.(type) {
	case sortDir:
		cmd.Log(v, "- Comparing dir:\n%v\n%v", x, y)
		return compareStringDelim(v, x, y, "/")
	case sortMod:
		cmd.Log(v, "- Comparing last modified times:\n%v\n%v", x, y)
		return compareTime(v, x, y)
	case sortName:
		cmd.Log(v, "- Comparing names:\n%v\n%v", x, y)
		return compareString(v, x, y)
	case sortSize:
		cmd.Log(v, "- Comparing sizes:\n%v\n%v", x, y)
		return compareInt(v, x, y)
	default:
		return false
	}
}

func compareStringDelim(v bool, x, y any, s string) bool {
	i := reflect.ValueOf(x).String()
	j := reflect.ValueOf(y).String()

	p := strings.Split(i, s)
	q := strings.Split(j, s)
	if len(p) > len(q) {
		cmd.Log(v, "- %s comes before %s", i, j)
		return true
	}
	if len(q) > len(p) {
		cmd.Log(v, "- %s comes before %s", j, i)
		return false
	}

	return compareString(v, x, y)
}

func compareTime(v bool, x, y any) bool {
	i := reflect.ValueOf(x).String()
	j := reflect.ValueOf(y).String()

	a, err := time.Parse(time.RFC3339Nano, i)
	if err != nil {
		cmd.Fatal("## " + err.Error())
	}
	b, err := time.Parse(time.RFC3339Nano, j)
	if err != nil {
		cmd.Fatal("## " + err.Error())
	}

	if a.After(b) {
		cmd.Log(v, "- %v comes before %v", a, b)
		return true
	}
	cmd.Log(v, "- %v comes before %v", b, a)
	return false
}

func compareString(v bool, x, y any) bool {
	i := reflect.ValueOf(x).String()
	j := reflect.ValueOf(y).String()

	n := strings.Compare(i, j)
	if n == 1 {
		cmd.Log(v, "- %s comes before %s", i, j)
		return true
	}
	cmd.Log(v, "- %s comes before %s", j, i)
	return false
}

func compareInt(v bool, x, y any) bool {
	i := reflect.ValueOf(x).Int()
	j := reflect.ValueOf(y).Int()

	if i > j {
		cmd.Log(v, "- %d comes before %d", i, j)
		return true
	}
	cmd.Log(v, "- %d comes before %d", j, i)
	return false
}
