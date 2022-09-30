package sort

import (
	"github.com/syhv-git/cookbook/cmd"
	"reflect"
	"strings"
	"time"
)

/*
All compare functions evaluate x > y in its own type context and returns the expression result

@devs-if you decide to add on to this project, handleDesc is the only function you need for handling your custom sort value types (you will need to add the handler and include it in the switch block)
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
		return compareDir(v, x, y)
	case sortMod:
		return compareMod(v, x, y)
	case sortName:
		return compareName(v, x, y)
	case sortSize:
		return compareSize(v, x, y)
	default:
		return false
	}
}

func compareDir(v bool, x, y any) bool {
	i := reflect.ValueOf(x).String()
	j := reflect.ValueOf(y).String()
	cmd.Log(v, "- Comparing dir:\n%s\n%s", i, j)

	p := strings.Split(i, "/")
	q := strings.Split(j, "/")
	if len(p) > len(q) {
		cmd.Log(v, "- %s has more parent directories", i)
		return true
	}

	n := strings.Compare(i, j)
	if n == 1 {
		cmd.Log(v, "- %s comes before %s", i, j)
		return true
	}
	cmd.Log(v, "- %s comes before %s", j, i)
	return false
}

func compareMod(v bool, x, y any) bool {
	i := reflect.ValueOf(x).String()
	j := reflect.ValueOf(y).String()
	cmd.Log(v, "- Comparing last modified times:\n%#v\n%#v", x, y)

	a, err := time.Parse(time.RFC3339Nano, i)
	if err != nil {
		cmd.Fatal("## " + err.Error())
	}
	b, err := time.Parse(time.RFC3339Nano, j)
	if err != nil {
		cmd.Fatal("## " + err.Error())
	}

	if a.After(b) {
		cmd.Log(v, "- %#v comes before %#v", a, b)
		return true
	}
	cmd.Log(v, "- %#v comes before %#v", b, a)
	return false
}

func compareName(v bool, x, y any) bool {
	i := reflect.ValueOf(x).String()
	j := reflect.ValueOf(y).String()
	cmd.Log(v, "- Comparing names: %s %s", i, j)

	n := strings.Compare(i, j)
	if n == 1 {
		cmd.Log(v, "- %s comes before %s", i, j)
		return true
	}
	cmd.Log(v, "- %s comes before %s", j, i)
	return false
}

func compareSize(v bool, x, y any) bool {
	i := reflect.ValueOf(x).Int()
	j := reflect.ValueOf(y).Int()
	cmd.Log(v, "- Comparing sizes: %d %d", i, j)

	if i > j {
		cmd.Log(v, "- %d comes before %d", i, j)
		return true
	}
	cmd.Log(v, "- %d comes before %d", j, i)
	return false
}
