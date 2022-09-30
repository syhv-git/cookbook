package sort

import (
	"github.com/syhv-git/cookbook/cmd"
	types "github.com/syhv-git/cookbook/file"
)

func BubbleSort(v, b bool, s string, t types.Tree) {
	cmd.Log(v, "*** Starting bubble sort")
	defer cmd.Log(v, "*** Ending bubble sort")

	switch s {
	case "dir":
		bubbleSorter(v, b, t, dir)
	case "mod":
		bubbleSorter(v, b, t, mod)
	case "name":
		bubbleSorter(v, b, t, name)
	case "size":
		bubbleSorter(v, b, t, size)
	default:
		cmd.Log(v, "* Unknown sort type: %s", s)
	}
}

func bubbleSorter[T constraint](v, b bool, t types.Tree, data T) {
	for i := 0; i < len(t); i++ {
		for j := 0; j < len(t)-i-1; j++ {
			if handleDesc(v, b, data.handle(t[j+1]), data.handle(t[j]), data) {
				t[j], t[j+1] = t[j+1], t[j]
			}
		}
	}
}
