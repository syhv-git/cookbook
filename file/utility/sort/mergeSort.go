package sort

import (
	"github.com/syhv-git/cookbook/cmd"
	types "github.com/syhv-git/cookbook/file"
)

func MergeSort(v, b bool, s string, t types.Tree) {
	cmd.Log(v, "*** Starting merge sort")
	defer cmd.Log(v, "*** Ending merge sort")

	switch s {
	case "dir":
		mergeSorter(v, b, t, 0, len(t)-1, dir)
	case "mod":
		mergeSorter(v, b, t, 0, len(t)-1, mod)
	case "name":
		mergeSorter(v, b, t, 0, len(t)-1, name)
	case "size":
		mergeSorter(v, b, t, 0, len(t)-1, size)
	default:
		cmd.Log(v, "* Unknown sort type: %s", s)
	}
}

func mergeSorter[T constraint](v, b bool, t types.Tree, l, r int, data T) {
	if l >= r {
		return
	}

	m := l + (r-l)/2
	mergeSorter(v, b, t, l, m, data)
	mergeSorter(v, b, t, m+1, r, data)
	merge(v, b, t, l, m, r, data)
}

func merge[T constraint](v, b bool, t types.Tree, l, m, r int, data T) {
	if handleDesc(v, b, data.handle(t[m]), data.handle(t[m+1]), data) {
		return
	}

	for m2 := m + 1; l <= m && m2 <= r; {
		if handleDesc(v, b, data.handle(t[m]), data.handle(t[m2]), data) {
			l++
			continue
		}
		i, n := m2, t[m2]
		for ; i > l; i-- {
			if handleDesc(v, b, data.handle(t[i-1]), data.handle(n), data) {
				n, t[i-1] = t[i-1], n
			}
			t[i] = t[i-1]
		}
		t[l] = n
		l, m, m2 = l+1, m+1, m2+1
	}
}
