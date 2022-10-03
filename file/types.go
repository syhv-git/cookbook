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

func NewNode(p string) (n Node, err error) {
	var i os.FileInfo
	p = path.Clean(p)
	i, err = os.Stat(p)
	if err != nil {
		return
	}
	return Node{Path: p, Nodr: i}, nil
}

type Tree []Node

func NewTree(src ...string) Tree {
	var t Tree
	for _, p := range src {
		n, err := NewNode(p)
		if err != nil {
			continue
		}
		t = t.Append(n)
	}
	return t
}

func (t Tree) Append(v ...Node) Tree {
	return append(t, v...)
}

// GetFilePaths returns all paths in the tree.
// offset defines the index to start the path for each string
func (t Tree) GetFilePaths(offset int) (s []string) {
	for _, x := range t {
		s = append(s, x.Path[offset:])
	}
	return
}
