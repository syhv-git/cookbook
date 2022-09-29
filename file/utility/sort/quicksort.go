package sort

import (
	"cookbook/cmd"
	. "cookbook/file"
	"reflect"
)

func QuickSort(v bool, t Tree, s string, b bool) {
	cmd.Log(v, "*** Starting quicksort")
	defer cmd.Log(v, "*** Ending quicksort")

	switch s {
	case "dir":
		sorter(v, t, b, dir)
	case "mod":
		sorter(v, t, b, mod)
	case "name":
		sorter(v, t, b, name)
	case "size":
		sorter(v, t, b, size)
	default:
		cmd.Log(v, "* Unknown sort type: %s", s)
	}
}

func sorter[T constraint](v bool, t Tree, b bool, data T) {
	if len(t) < 2 {
		return
	}
	p := partition(v, t, b, data)
	sorter(v, t[:p], b, data)
	sorter(v, t[p+1:], b, data)
}

func partition[T constraint](v bool, t Tree, b bool, data T) int {
	end := len(t) - 1
	pivot := data.handle(t[end])
	i := -1
	for j, n := range t {
		h := data.handle(n)
		if b {
			if compare(v, h, pivot, data) {
				i++
				t[i], t[j] = t[j], t[i]
			}
			continue
		}
		if compare(v, pivot, h, data) {
			i++
			t[i], t[j] = t[j], t[i]
		}
	}
	i++
	t[i], t[end] = t[end], t[i]
	return i
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
