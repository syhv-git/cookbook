package utility

import (
	. "cookbook/file"
	"path"
)

type constraint interface {
	string | int64
}

var (
	dir  = func(n Node) string { return path.Base(n.Path) }
	mod  = func(n Node) string { return n.Info.ModTime().String() }
	name = func(n Node) string { return n.Info.Name() }
	size = func(n Node) int64 { return n.Info.Size() }
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

func sorter[T constraint](t Tree, b bool, start, end int, data func(n Node) T) {
	if start < end {
		p := partition(t, b, start, end, data)
		sorter(t, b, start, p, data)
		sorter(t, b, p+1, end, data)
	}
}

func partition[T constraint](t Tree, b bool, start, end int, data func(n Node) T) int {
	pivot := data(t[end])
	i := start - 1
	for j := start; j <= end; j++ {
		n := t[j]
		v := data(n)
		if b {
			if v > pivot {
				i++
				tmp := t[j]
				t[j] = t[i]
				t[i] = tmp
			}
			continue
		}
		if v < pivot {
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
