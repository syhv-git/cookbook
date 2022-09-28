package tests

import (
	types "cookbook/file"
	"testing"
)

var no, _ = types.NewNode("")

func TestNewNode(t *testing.T) {
	node, err := types.NewNode("../run.go")
	if err != nil {
		t.Fatal("## " + err.Error())
	}
	if node == no {
		t.Error("## Error when creating a new node")
	}
}

func TestNewTree(t *testing.T) {
	tree := types.NewTree(".")
	if len(tree) < 1 {
		t.Error("## Error when create a new tree")
	}
}

func TestAppend(t *testing.T) {
	var tree types.Tree
	n1, _ := types.NewNode("..")
	tree = tree.Append(n1)
	if len(tree) != 1 {
		t.Error("## Error when appending a node to a tree")
	}
}

func TestGetPaths(t *testing.T) {
	tree := types.NewTree(".")
	s := tree.GetPaths(0)
	if len(s) < 1 {
		t.Error("## Error when getting paths")
	}
}
