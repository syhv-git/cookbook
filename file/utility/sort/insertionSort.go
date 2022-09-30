package sort

import (
	"github.com/syhv-git/cookbook/cmd"
	types "github.com/syhv-git/cookbook/file"
)

func InsertionSort(v, b bool, s string, t types.Tree) {
	cmd.Log(v, "*** Starting insertion sort")
	defer cmd.Log(v, "*** Ending insertion sort")

	switch s {
	case "dir":
		insertionSorter(v, b, t, dir)
	case "mod":
		insertionSorter(v, b, t, mod)
	case "name":
		insertionSorter(v, b, t, name)
	case "size":
		insertionSorter(v, b, t, size)
	default:
		cmd.Log(v, "* Unknown sort type: %s", s)
	}
}

func insertionSorter[T constraint](v, b bool, t types.Tree, data T) {
	for i := 1; i < len(t); i++ {
		for j := i; j > 0 && handleDesc(v, b, data.handle(t[j]), data.handle(t[j-1]), data); j-- {
			t[j-1], t[j] = t[j], t[j-1]
		}
	}
}
