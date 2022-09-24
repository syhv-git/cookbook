package forensics

import (
	. "cookbook/file"
	"log"
	"os"
)

func dirWalker(list Tree) (all Tree) {
	c := make(chan Node)
	e := make(chan error)
	go walk(list, c, e)

wait:
	for {
		select {
		case x, ok := <-c:
			if !ok {
				break wait
			}
			all.Append(x)
		case err := <-e:
			log.Fatal(err.Error())
		}
	}
	return
}

func walk(files Tree, c chan Node, e chan error) {
	defer close(c)
	if len(files) < 1 {
		return
	}

	for _, f := range files {
		if !f.Info.IsDir() {
			c <- f
			continue
		}
		cont, _ := os.ReadDir(f.Path)
		sub := handleContents(f.Path, cont)

		c2 := make(chan Node)
		e2 := make(chan error)
		go walk(sub, c2, e2)

	wait2:
		for {
			select {
			case x, ok := <-c2:
				if !ok {
					break wait2
				}
				c <- x
			case err := <-e2:
				e <- err
				return
			}
		}
	}
}

func handleContents(p string, c []os.DirEntry) (res Tree) {
	var err error
	for _, de := range c {
		t := Node{
			Path: p + "/" + de.Name(),
		}
		t.Info, err = os.Stat(t.Path)
		if err != nil {
			continue
		}
		res.Append(t)
	}
	return
}

// Enumerate recursively walks through directories and sorts the discovered files before returning the list.
// sort must be one of ("dir" | "mod" | "name" | "size")
// desc defines whether the contents are sorted in descending order or ascending order
// paths is a variadic list of paths to enumerate
func Enumerate(sort string, desc bool, paths ...string) (res Tree) {
	for _, p := range paths {
		n, err := NewNode(p)
		if err != nil {
			log.Println(err)
			continue
		}
		res.Append(n)
	}

	res = dirWalker(res)
	res.Sort(sort, desc)
	return
}

func Extract(dest string, paths ...string) (res []byte) {
	var ret Tree
	for _, p := range paths {
		n, err := NewNode(p)
		if err != nil {
			log.Fatal(err.Error())
		}
		ret.Append(n)
	}
	ret = dirWalker(ret)

	return
}
