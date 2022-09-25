package utility

import (
	. "cookbook/file"
)

var (
	dir  sortDir
	mod  sortMod
	name sortName
	size sortSize
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
			if data.compare(v, pivot) {
				i++
				t[i], t[j] = t[j], t[i]
			}
			continue
		}
		if data.compare(pivot, v) {
			i++
			t[i], t[j] = t[j], t[i]
		}
	}
	i++
	t[i], t[end] = t[end], t[i]
	return i
}
