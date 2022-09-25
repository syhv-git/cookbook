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
	pivot := data.handle(t[end])
	i := start - 1
	for j := start; j <= end; j++ {
		n := t[j]
		v := data.handle(n)
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
