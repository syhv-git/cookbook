package tests

import (
	types "github.com/syhv-git/cookbook/file"
	"github.com/syhv-git/cookbook/file/utility/sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	f1, f2 := "../.gitattributes", "../file/forensics/enumerate.go"
	tree := types.NewTree(f1, f2)

	sort.QuickSort(true, tree, "dir", true)
	if tree[0].Path != f2 {
		t.Errorf("## Error when sorting a tree %#v", tree)
	}

	sort.QuickSort(true, tree, "mod", true)
	if tree[0].Path != f2 {
		t.Errorf("## Error when sorting a tree %#v", tree)
	}

	sort.QuickSort(true, tree, "name", true)
	if tree[0].Path != f2 {
		t.Errorf("## Error when sorting a tree %#v", tree)
	}

	sort.QuickSort(true, tree, "size", true)
	if tree[0].Path != f2 {
		t.Errorf("## Error when sorting a tree %#v", tree)
	}
}

func TestInsertionSort(t *testing.T) {
	f1, f2 := "../.gitattributes", "../file/forensics/enumerate.go"
	tree := types.NewTree(f1, f2)

	sort.InsertionSort(true, tree, "dir", true)
	if tree[0].Path != f2 {
		t.Errorf("## Error when sorting a tree %#v", tree)
	}

	//sort.InsertionSort(true, tree, "mod", true)
	//if tree[0].Path != f2 {
	//t.Errorf("## Error when sorting a tree %#v", tree)
	//}
	//
	//sort.InsertionSort(true, tree, "name", true)
	//if tree[0].Path != f2 {
	//t.Errorf("## Error when sorting a tree %#v", tree)
	//}
	//
	//sort.InsertionSort(true, tree, "size", true)
	//if tree[0].Path != f2 {
	//t.Errorf("## Error when sorting a tree %#v", tree)
	//}
}
