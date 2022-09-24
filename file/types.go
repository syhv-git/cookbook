package types

import (
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
