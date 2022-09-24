package tests

import (
	types "cookbook/file"
	"cookbook/file/utility"
	"os"
	"testing"
)

func TestQuickSort(t *testing.T) {
	var tree types.Tree
	f1, f2 := "../go.mod", "../file/forensics/enumerate.go"
	info1, err := os.Stat(f1)
	if err != nil {
		t.Fatal(err.Error())
	}
	info2, err := os.Stat(f2)
	if err != nil {
		t.Fatal(err.Error())
	}
	tree = tree.Append(types.Node{Path: f1, Info: info1}, types.Node{Path: f2, Info: info2})
	utility.QuickSort(tree, "size", true)
	if tree[0].Path != f2 {
		t.Error("Error when sorting a tree")
	}
}
