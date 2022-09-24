package types

import (
	"cookbook/file/utility"
	"os"
)

type Node struct {
	Path string
	Info os.FileInfo
}

func NewNode(path string) (Node, error) {
	i, err := os.Stat(path)
	if err != nil {
		return Node{}, err
	}
	return Node{Path: path, Info: i}, nil
}

type Tree []Node

func (t Tree) Append(v ...Node) Tree {
	return append(t, v...)
}

func (t Tree) Sort(s string, b bool) {
	utility.QuickSort(t, s, b)
}
