package sort

import (
	"github.com/syhv-git/cookbook/cmd"
	types "github.com/syhv-git/cookbook/file"
)

func SelectionSort(v, b bool, s string, t types.Tree) {
	cmd.Log(v, "*** Starting selection sort")
	defer cmd.Log(v, "*** Ending selection sort")

	switch s {
	case "dir":
		selectionSorter(v, b, t, dir)
	case "mod":
		selectionSorter(v, b, t, mod)
	case "name":
		selectionSorter(v, b, t, name)
	case "size":
		selectionSorter(v, b, t, size)
	default:
		cmd.Log(v, "* Unknown sort type: %s", s)
	}
}

func selectionSorter[T constraint](v, b bool, t types.Tree, data T) {
	for i := 0; i < len(t); i++ {
		m := i
		for j := i + 1; j < len(t); j++ {
			if handleDesc(v, b, data.handle(t[j]), data.handle(t[m]), data) {
				m = j
			}
		}
		t[i], t[m] = t[m], t[i]
	}
}
