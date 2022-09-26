package utility

import (
	. "cookbook/file"
	"time"
)

type sortDir string

type sortMod string

type sortName string

type sortSize int64

type constraint interface {
	sortDir | sortMod | sortName | sortSize
	handle(Node) any
}

func (sortDir) handle(n Node) any  { return n.Path }
func (sortMod) handle(n Node) any  { return n.ModTime().Format(time.RFC3339Nano) }
func (sortName) handle(n Node) any { return n.Name() }
func (sortSize) handle(n Node) any { return n.Size() }

var (
	dir  sortDir
	mod  sortMod
	name sortName
	size sortSize
)
