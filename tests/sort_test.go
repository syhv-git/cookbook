package tests

import (
	types "github.com/syhv-git/cookbook/file"
	"github.com/syhv-git/cookbook/file/utility/sort"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	f1, f2 := "../.gitattributes", "../file/forensics/enumerate.go"
	tree := types.NewTree(f1, f2)

	sort.InsertionSort(true, true, "dir", tree)
	if tree[0].Path != f2 {
		t.Errorf("## Error when sorting a tree %#v", tree)
	}

	sort.InsertionSort(true, true, "mod", tree)
	if tree[0].Path != f2 {
		t.Errorf("## Error when sorting a tree %#v", tree)
	}

	sort.InsertionSort(true, true, "name", tree)
	if tree[0].Path != f2 {
		t.Errorf("## Error when sorting a tree %#v", tree)
	}

	sort.InsertionSort(true, true, "size", tree)
	if tree[0].Path != f2 {
		t.Errorf("## Error when sorting a tree %#v", tree)
	}
}

func TestQuickSort(t *testing.T) {
	f1, f2 := "../.gitattributes", "../file/forensics/enumerate.go"
	tree := types.NewTree(f1, f2)

	sort.QuickSort(true, true, "dir", tree)
	if tree[0].Path != f2 {
		t.Errorf("## Error when sorting a tree %#v", tree)
	}

	sort.QuickSort(true, true, "mod", tree)
	if tree[0].Path != f2 {
		t.Errorf("## Error when sorting a tree %#v", tree)
	}

	sort.QuickSort(true, true, "name", tree)
	if tree[0].Path != f2 {
		t.Errorf("## Error when sorting a tree %#v", tree)
	}

	sort.QuickSort(true, true, "size", tree)
	if tree[0].Path != f2 {
		t.Errorf("## Error when sorting a tree %#v", tree)
	}
}

func TestSelectionSort(t *testing.T) {
	f1, f2 := "../.gitattributes", "../file/forensics/enumerate.go"
	tree := types.NewTree(f1, f2)

	sort.SelectionSort(true, true, "dir", tree)
	if tree[0].Path != f2 {
		t.Errorf("## Error when sorting a tree %#v", tree)
	}

	sort.SelectionSort(true, true, "mod", tree)
	if tree[0].Path != f2 {
		t.Errorf("## Error when sorting a tree %#v", tree)
	}

	sort.SelectionSort(true, true, "name", tree)
	if tree[0].Path != f2 {
		t.Errorf("## Error when sorting a tree %#v", tree)
	}

	sort.SelectionSort(true, true, "size", tree)
	if tree[0].Path != f2 {
		t.Errorf("## Error when sorting a tree %#v", tree)
	}
}
