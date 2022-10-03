package tests

import (
	types "github.com/syhv-git/cookbook/file"
	"github.com/syhv-git/cookbook/file/utility/sort"
	"testing"
)

func TestSortingFunction(t *testing.T) {
	f1, f2, f3, f4 := "../.gitattributes", "../file/forensics/enumerate.go", "enumerate_test.go", "../file/utility/image.go"
	tree := types.NewTree(f1, f2, f3, f4)

	//sort.BubbleSort(true, true, "name", tree)
	//sort.InsertionSort(true, true, "size", tree)
	sort.MergeSort(true, false, "mod", tree)

	if tree[0].Path != f1 {
		t.Errorf("## Error when sorting function; %#v", tree)
	}
	tree[0], tree[1], tree[2], tree[3] = tree[2], tree[0], tree[3], tree[1]

	sort.QuickSort(true, true, "dir", tree)
	//sort.SelectionSort(true, true, "name", tree)

	if tree[0].Path != f4 {
		t.Errorf("## Error when sorting function; %#v", tree)
	}
}
