package sort

import (
	. "cookbook/file"
	"time"
)

type sortDir string

type sortMod string

type sortName string

type sortSize int64

func (sortDir) handle(n Node) any  { return n.Path }
func (sortMod) handle(n Node) any  { return n.ModTime().Format(time.RFC3339Nano) }
func (sortName) handle(n Node) any { return n.Name() }
func (sortSize) handle(n Node) any { return n.Size() }

type constraint interface {
	sortDir | sortMod | sortName | sortSize
	handle(Node) any
}

var (
	dir  sortDir
	mod  sortMod
	name sortName
	size sortSize
)
