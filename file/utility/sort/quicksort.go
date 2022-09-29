package sort

import (
	"github.com/syhv-git/cookbook/cmd"
	. "github.com/syhv-git/cookbook/file"
)

// QuickSort sorts t with quicksort based on the FileInfo type on s. b determines sortDescending ? true : false
func QuickSort(v bool, t Tree, s string, b bool) {
	cmd.Log(v, "*** Starting quicksort")
	defer cmd.Log(v, "*** Ending quicksort")

	switch s {
	case "dir":
		quickSorter(v, t, b, dir)
	case "mod":
		quickSorter(v, t, b, mod)
	case "name":
		quickSorter(v, t, b, name)
	case "size":
		quickSorter(v, t, b, size)
	default:
		cmd.Log(v, "* Unknown sort type: %s", s)
	}
}

func quickSorter[T constraint](v bool, t Tree, b bool, data T) {
	if len(t) < 2 {
		return
	}
	p := partition(v, t, b, data)
	quickSorter(v, t[:p], b, data)
	quickSorter(v, t[p+1:], b, data)
}

func partition[T constraint](v bool, t Tree, b bool, data T) int {
	end := len(t) - 1
	pivot := data.handle(t[end])
	i := -1
	for j, n := range t {
		if handleDesc(v, b, data.handle(n), pivot, data) {
			i++
			t[i], t[j] = t[j], t[i]
		}
	}
	i++
	t[i], t[end] = t[end], t[i]
	return i
}
