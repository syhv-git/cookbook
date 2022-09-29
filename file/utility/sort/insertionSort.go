package sort

import (
	"github.com/syhv-git/cookbook/cmd"
	. "github.com/syhv-git/cookbook/file"
)

func InsertionSort(v bool, t Tree, s string, b bool) {
	cmd.Log(v, "*** Starting insertion sort")
	defer cmd.Log(v, "*** Ending insertion sort")

	switch s {
	case "dir":
		insertionSorter(v, t, b, dir)
	case "mod":
		insertionSorter(v, t, b, mod)
	case "name":
		insertionSorter(v, t, b, name)
	case "size":
		insertionSorter(v, t, b, size)
	default:
		cmd.Log(v, "* Unknown sort type: %s", s)
	}
}

func insertionSorter[T constraint](v bool, t Tree, b bool, data T) {
	for i, j := 1, 0; i < len(t); i++ {
		for j >= 0 && handleDesc(v, b, data.handle(t[j]), data.handle(t[i]), data) {
			t[j+1] = t[j]
			j -= 1
		}
		t[j+1] = t[i]
	}
}
