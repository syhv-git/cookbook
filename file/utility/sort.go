package utility

import (
	. "cookbook/file"
	"log"
	"path"
	"reflect"
	"strings"
	"time"
)

var (
	dir  = sortDir(func(n Node) string { return path.Dir(n.Path) })
	mod  = sortMod(func(n Node) time.Time { return n.Info.ModTime() })
	name = sortName(func(n Node) string { return n.Info.Name() })
	size = sortSize(func(n Node) int64 { return n.Info.Size() })
)

type constraint interface {
	sortDir | sortMod | sortName | sortSize
	compare(x, y any) bool
}

type sortDir func(f func(n Node) string) string
type sortMod func(f func(n Node) time.Time) time.Time
type sortName func(f func(n Node) string) string
type sortSize func(f func(n Node) int64) int64

func (d sortDir) compare(x, y any) bool {
	i := reflect.ValueOf(x).String()
	j := reflect.ValueOf(y).String()
	n := strings.Compare(i, j)
	if n == 1 {
		return true
	}
	return false
}

func (d sortMod) compare(x, y any) bool {
	i := reflect.ValueOf(x).String()
	j := reflect.ValueOf(y).String()
	a, err := time.Parse(time.Layout, i)
	if err != nil {
		log.Fatal(err.Error())
	}
	b, err := time.Parse(time.Layout, j)
	if err != nil {
		log.Fatal(err.Error())
	}
	if a.After(b) {
		return true
	}
	return false
}

func (d sortName) compare(x, y any) bool {
	i := reflect.ValueOf(x).String()
	j := reflect.ValueOf(y).String()
	n := strings.Compare(i, j)
	if n == 1 {
		return true
	}
	return false
}

func (d sortSize) compare(x, y any) bool {
	i := reflect.ValueOf(x).Int()
	j := reflect.ValueOf(y).Int()
	if i > j {
		return true
	}
	return false
}

func QuickSort(t Tree, s string, b bool) {
	switch s {
	case "dir":
		sorter(t, b, 0, len(t)-1, dir)
	case "mod":
		sorter(t, b, 0, len(t)-1, mod)
	case "name":
		sorter(t, b, 0, len(t)-1, name)
	case "size":
		sorter(t, b, 0, len(t)-1, size)
	}
}

func sorter[T constraint](t Tree, b bool, start, end int, data T) {
	if start < end {
		p := partition(t, b, start, end, data)
		sorter(t, b, start, p-1, data)
		sorter(t, b, p+1, end, data)
	}
}

func partition[T constraint](t Tree, b bool, start, end int, data T) int {
	pivot := data(t[end])
	i := start - 1
	for j := start; j <= end; j++ {
		n := t[j]
		v := data(n)
		if b {
			if data.compare(v, pivot) {
				i++
				tmp := t[j]
				t[j] = t[i]
				t[i] = tmp
			}
			continue
		}
		if data.compare(pivot, v) {
			i++
			tmp := t[j]
			t[j] = t[i]
			t[i] = tmp
		}
	}
	i++
	tmp := t[end]
	t[end] = t[i]
	t[i] = tmp
	return i
}
