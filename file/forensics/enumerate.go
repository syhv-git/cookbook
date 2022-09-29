package forensics

import (
	"bytes"
	"cookbook/cmd"
	. "cookbook/file"
	"cookbook/file/utility/sort"
	"io"
	"os"
	"path"
)

// Enumerate recursively walks through directories and sorts the discovered files before returning the list.
// sort must be one of ("dir" | "mod" | "name" | "size")
// desc defines whether the contents are sorted in descending order or ascending order
// paths is a variadic list of paths to enumerate
func Enumerate(v bool, sortBy string, desc bool, paths ...string) Tree {
	res := NewTree(paths...)
	res = dirWalker(v, res)
	sort.QuickSort(v, res, sortBy, desc)
	return res
}

// Extractor extracts the files at src and concatenates the contents into the writer.
func Extractor(v bool, dst io.Writer, src ...string) {
	res := NewTree(src...)
	res = dirWalker(v, res)
	if err := extractor(v, dst, res); err != nil {
		cmd.Fatal("## " + err.Error())
	}
}

// ExtractCopy extracts the files at src and creates a new file dest with the concatenated contents.
func ExtractCopy(v bool, dst string, src ...string) {
	res := NewTree(src...)
	res = dirWalker(v, res)

	buf := bytes.NewBuffer(nil)
	if err := extractor(v, buf, res); err != nil {
		cmd.Fatal("## " + err.Error())
	}

	cmd.Log(v, "- Creating output file: %s", dst)
	if err := os.MkdirAll(path.Dir(dst), 0777); err != nil && !os.IsExist(err) {
		cmd.Fatal("## " + err.Error())
	}
	if err := os.WriteFile(dst, buf.Bytes(), 0666); err != nil {
		cmd.Fatal("## " + err.Error())
	}
	cmd.Log(v, "* Successfully copied contents from %v", res)
}

func extractor(v bool, w io.Writer, t Tree) error {
	cmd.Log(v, "*** Starting extraction")
	defer cmd.Log(v, "*** Ending extraction")

	for _, x := range t {
		cmd.Log(v, "- Extracting contents: %s", x.Path)
		f, err := os.ReadFile(x.Path)
		if err != nil {
			continue
		}
		if err = os.MkdirAll(path.Dir(x.Path), 0777); err != nil && !os.IsExist(err) {
			cmd.Fatal("## " + err.Error())
		}
		if _, err = w.Write(f); err != nil {
			return err
		}
	}
	return nil
}

func dirWalker(v bool, list Tree) Tree {
	var all Tree
	cmd.Log(v, "*** Starting enumeration")
	defer cmd.Log(v, "*** Ending enumeration")

	c := make(chan Node)
	e := make(chan error)
	go walk(v, list, c, e)

wait:
	for {
		select {
		case x, ok := <-c:
			if !ok {
				break wait
			}
			all = all.Append(x)
		case err := <-e:
			cmd.Fatal("## " + err.Error())
		}
	}
	return all
}

func walk(v bool, files Tree, c chan Node, e chan error) {
	defer close(c)
	if len(files) < 1 {
		return
	}

	for _, f := range files {
		if !f.IsDir() {
			c <- f
			continue
		}
		cont, _ := os.ReadDir(f.Path)
		sub := handleContents(f.Path, cont)

		c2 := make(chan Node)
		e2 := make(chan error)
		go walk(v, sub, c2, e2)

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
	for _, de := range c {
		t, err := NewNode(p + "/" + de.Name())
		if err != nil {
			continue
		}
		res = res.Append(t)
	}
	return
}
