package utility

import (
	. "cookbook/file"
	"path"
	"time"
)

type constraint interface {
	string | int64 | time.Time
}

var (
	dir  = func(n Node) string { return path.Base(n.Path) }
	mod  = func(n Node) time.Time { return n.Info.ModTime() }
	name = func(n Node) string { return n.Info.Name() }
	size = func(n Node) int64 { return n.Info.Size() }
)

func QuickSort(t Tree, s string, b bool) {
	switch s {
	case "dir":
		sorter(t, b, 0, len(t), dir)
	case "mod":
		sorter(t, b, 0, len(t), mod)
	case "name":
		sorter(t, b, 0, len(t), name)
	case "size":
		sorter(t, b, 0, len(t), size)
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
	return 0
}
