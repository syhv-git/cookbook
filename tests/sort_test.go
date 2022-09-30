package tests

import (
	types "github.com/syhv-git/cookbook/file"
	"github.com/syhv-git/cookbook/file/utility/sort"
	"testing"
)

func TestSortingFunction(t *testing.T) {
	f1, f2, f3 := "../.gitattributes", "../file/forensics/enumerate.go", "enumerate_test.go"
	tree := types.NewTree(f1, f2, f3)

	//sort.BubbleSort(true, true, "name", tree)
	//sort.InsertionSort(true, true, "size", tree)
	sort.MergeSort(true, true, "mod", tree)

	if tree[0].Path != f2 {
		t.Errorf("## Error when sorting function; %#v", tree)
	}
	tree[0], tree[1], tree[2] = tree[2], tree[0], tree[1]

	sort.QuickSort(true, true, "dir", tree)
	//sort.SelectionSort(true, true, "name", tree)

	if tree[0].Path != f2 {
		t.Errorf("## Error when sorting function; %#v", tree)
	}
}
