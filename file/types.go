package types

import (
	"os"
	"path"
)

type Nodr interface {
	os.FileInfo
}

type Node struct {
	Path string
	Nodr
}

func NewNode(p string) (Node, error) {
	p = path.Clean(p)
	i, err := os.Stat(p)
	if err != nil {
		return Node{}, err
	}
	return Node{Path: p, Nodr: i}, nil
}

type Tree []Node

func (t Tree) Append(v ...Node) Tree {
	return append(t, v...)
}
