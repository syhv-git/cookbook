package sort

import (
	"github.com/syhv-git/cookbook/cmd"
	types "github.com/syhv-git/cookbook/file"
)

// QuickSort performs a quicksort on t with b defining the vertical sort order and s defining the sort value
func QuickSort(v, b bool, s string, t types.Tree) {
	cmd.Log(v, "*** Starting quicksort")
	defer cmd.Log(v, "*** Ending quicksort")

	switch s {
	case "dir":
		quickSorter(v, b, t, dir)
	case "mod":
		quickSorter(v, b, t, mod)
	case "name":
		quickSorter(v, b, t, name)
	case "size":
		quickSorter(v, b, t, size)
	default:
		cmd.Log(v, "* Unknown sort type: %s", s)
	}
}

func quickSorter[T constraint](v, b bool, t types.Tree, data T) {
	if len(t) < 2 {
		return
	}
	p := partition(v, b, t, data)
	quickSorter(v, b, t[:p], data)
	quickSorter(v, b, t[p+1:], data)
}

func partition[T constraint](v, b bool, t types.Tree, data T) int {
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
