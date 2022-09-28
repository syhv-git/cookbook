package tests

import (
	types "cookbook/file"
	"cookbook/file/utility/sort"
	"os"
	"testing"
)

func TestQuickSort(t *testing.T) {
	var tree types.Tree
	f1, f2 := "../.gitattributes", "../file/forensics/enumerate.go"
	info1, err := os.Stat(f1)
	if err != nil {
		t.Fatal("## " + err.Error())
	}
	info2, err := os.Stat(f2)
	if err != nil {
		t.Fatal("## " + err.Error())
	}

	tree = tree.Append(types.Node{Path: f1, Nodr: info1}, types.Node{Path: f2, Nodr: info2})

	sort.QuickSort(true, tree, "dir", true)
	if tree[0].Path != f2 {
		t.Error("## Error when sorting a tree")
	}

	sort.QuickSort(true, tree, "mod", true)
	if tree[0].Path != f2 {
		t.Error("Error when sorting a tree")
	}

	sort.QuickSort(true, tree, "name", true)
	if tree[0].Path != f2 {
		t.Error("Error when sorting a tree")
	}

	sort.QuickSort(true, tree, "size", true)
	if tree[0].Path != f2 {
		t.Error("Error when sorting a tree")
	}
}
