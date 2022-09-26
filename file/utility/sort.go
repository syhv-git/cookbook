package utility

import (
	. "cookbook/file"
	"reflect"
)

func QuickSort(t Tree, s string, b bool) {
	switch s {
	case "dir":
		sorter(t, b, dir)
	case "mod":
		sorter(t, b, mod)
	case "name":
		sorter(t, b, name)
	case "size":
		sorter(t, b, size)
	}
}

func sorter[T constraint](t Tree, b bool, data T) {
	if len(t) < 2 {
		return
	}
	p := partition(t, b, data)
	sorter(t[:p], b, data)
	sorter(t[p+1:], b, data)
}

func partition[T constraint](t Tree, b bool, data T) int {
	end := len(t) - 1
	pivot := data.handle(t[end])
	i := -1
	for j, n := range t {
		v := data.handle(n)
		if b {
			if compare(v, pivot, data) {
				i++
				t[i], t[j] = t[j], t[i]
			}
			continue
		}
		if compare(pivot, v, data) {
			i++
			t[i], t[j] = t[j], t[i]
		}
	}
	i++
	t[i], t[end] = t[end], t[i]
	return i
}

func compare[K any, T constraint](x, y K, data T) bool {
	t := reflect.ValueOf(data)
	d := reflect.Indirect(t)
	i := d.Interface()
	switch i.(type) {
	case sortDir:
		return compareDir(x, y)
	case sortMod:
		return compareMod(x, y)
	case sortName:
		return compareName(x, y)
	case sortSize:
		return compareSize(x, y)
	default:
		return false
	}
}
