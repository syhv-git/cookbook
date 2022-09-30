package utility

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"github.com/syhv-git/cookbook/cmd"
	types "github.com/syhv-git/cookbook/file"
	"github.com/syhv-git/cookbook/file/forensics"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// CompressNew compresses the files at src into dst based on method.
func CompressNew(v bool, dst string, src ...string) {
	cmd.Log(v, "*** Starting compression")
	defer cmd.Log(v, "*** Ending compression")
	if path.IsAbs(dst) {
		cmd.Fatal("## Destination %s must be relative", dst)
	}

	cmd.Log(v, "- Creating destination file: %s", dst)
	if err := os.MkdirAll(path.Dir(path.Clean(dst)), 0777); err != nil && !os.IsExist(err) {
		cmd.Fatal("## " + err.Error())
	}
	f, err := os.OpenFile(dst, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {
		cmd.Fatal("## " + err.Error())
	}
	defer f.Close()
	info, _ := f.Stat()
	if info.IsDir() {
		cmd.Fatal("## Destination %s cannot be a directory", dst)
	}

	n := strings.Split(path.Base(dst), ".")
	cmd.Log(v, "- Creating temp dir")
	tmp, err := os.MkdirTemp(path.Dir(dst), n[0])
	if err != nil {
		cmd.Fatal("## " + err.Error())
	}
	defer os.RemoveAll(tmp)

	switch filepath.Ext(path.Base(dst)) {
	case ".gz":
		compressGZ(v, tmp, n[0]+"/", f, src)
	case ".tar":
		compressTar(v, tmp, n[0]+"/", f, src)
	case ".zip":
		compressZip(v, tmp, n[0]+"/", f, src)
	default:
		cmd.Log(v, "* Unknown compression method: %#v\n- Attempting Zip compression", strings.Join(n[1:], "."))
		compressZip(v, tmp, n[0]+"/", f, src)
	}
}

func compressTar(v bool, tmp, dst string, w io.Writer, src []string) {
	cmd.Log(v, "*** Starting Tar compression")
	defer cmd.Log(v, "*** Ending Tar compression")

	for _, x := range src {
		forensics.ExtractCopy(v, path.Join(tmp, path.Base(x)), x)
	}
	a := forensics.Enumerate(v, "", "", false, tmp)

	tw := tar.NewWriter(w)
	defer tw.Close()
	archive(dst, a, tw)
}

func archive(dst string, t types.Tree, tw *tar.Writer) {
	for _, x := range t {
		p := strings.Split(x.Path, "/")
		h := &tar.Header{
			Name:    dst + strings.Join(p[1:], "/"),
			Mode:    int64(x.Mode()),
			ModTime: x.ModTime(),
			Size:    x.Size(),
		}

		if err := tw.WriteHeader(h); err != nil {
			cmd.Fatal("## " + err.Error())
		}
		b, _ := os.ReadFile(x.Path)
		if _, err := tw.Write(b); err != nil {
			cmd.Fatal("## " + err.Error())
		}
	}
}

func compressGZ(v bool, tmp, dst string, w io.Writer, src []string) {
	cmd.Log(v, "*** Starting GZip compression")
	defer cmd.Log(v, "*** Ending GZip compression")

	for _, x := range src {
		n := path.Clean(strings.Join(strings.Split(x, ".."), ""))
		forensics.ExtractCopy(v, path.Join(tmp, n), x)
	}
	a := forensics.Enumerate(v, "", "", false, tmp)

	g := gzip.NewWriter(w)
	tw := tar.NewWriter(g)
	defer g.Close()
	defer tw.Close()
	archive(dst, a, tw)
}

func compressZip(v bool, tmp, dst string, w io.Writer, src []string) {
	cmd.Log(v, "*** Starting Zip compression")
	defer cmd.Log(v, "*** Ending Zip compression")

	for _, x := range src {
		n := path.Clean(strings.Join(strings.Split(x, ".."), ""))
		forensics.ExtractCopy(v, path.Join(tmp, n), x)
	}
	a := forensics.Enumerate(v, "", "", false, tmp)

	z := zip.NewWriter(w)
	defer z.Close()
	for _, p := range a {
		f, err := z.Create(dst + p.Path[len(tmp)-1:])
		if err != nil {
			cmd.Fatal("## " + err.Error())
		}
		buf := bytes.NewBuffer(nil)
		b, _ := os.ReadFile(p.Path)
		buf.Write(b)
		if _, err = f.Write(buf.Bytes()); err != nil {
			cmd.Fatal("## " + err.Error())
		}
	}
}

// Decompress decompresses an archive or compressed file and retrieves the contents.
// dst must be a directory. If dst is a nil string then the path from the source file is used.
func Decompress(v bool, dst, src string) {
	if dst != "" {
		if err := os.MkdirAll(dst, 0777); err != nil && !os.IsExist(err) {
			cmd.Fatal("## " + err.Error())
		}
	} else {
		cmd.Fatal("## No destination file provided")
	}
	f, err := os.Open(src)
	if err != nil {
		cmd.Fatal("## " + err.Error())
	}

	switch filepath.Ext(path.Base(src)) {
	case ".tar":
		decompressTar(v, f, dst)
	case ".gz":
		decompressGZ(v, f, dst)
	case ".zip":
		decompressZip(v, dst, src)
	default:
		cmd.Log(v, "* Unknown compression method: %#v\n- Attempting Zip decompression", strings.Join(strings.Split(path.Base(src), ".")[1:], "."))
		decompressZip(v, dst, src)
	}
}

func decompressTar(v bool, r io.Reader, dst string) {
	cmd.Log(v, "*** Starting Tar decompression")
	defer cmd.Log(v, "*** Ending Tar decompression")

	tr := tar.NewReader(r)
	unarchive(v, dst, tr)
}

func unarchive(v bool, dst string, tr *tar.Reader) {
	for {
		var (
			f   *os.File
			err error
		)

		h, e := tr.Next()
		if e == io.EOF {
			break
		}
		cmd.Log(v, "- Extracting: %s", h.Name)
		if h.Typeflag == tar.TypeDir {
			if err = os.MkdirAll(path.Join(dst, h.Name), 0777); err != nil && !os.IsExist(err) {
				cmd.Fatal("## " + err.Error())
			}
			continue
		}

		if err = os.MkdirAll(path.Join(dst, path.Dir(h.Name)), 0777); err != nil && !os.IsExist(err) {
			cmd.Fatal("## " + err.Error())
		}
		f, err = os.Create(path.Join(dst, h.Name))
		if err != nil {
			cmd.Fatal("## " + err.Error())
		}
		if _, err = io.Copy(f, tr); err != nil {
			f.Close()
			cmd.Fatal("## " + err.Error())
		}
		f.Close()
	}
}

func decompressGZ(v bool, r io.Reader, dst string) {
	cmd.Log(v, "*** Starting GZip decompression")
	defer cmd.Log(v, "*** Ending GZip decompression")

	gr, err := gzip.NewReader(r)
	defer gr.Close()
	if err != nil {
		cmd.Fatal("## " + err.Error())
	}
	tr := tar.NewReader(gr)
	unarchive(v, dst, tr)
}

func decompressZip(v bool, dst, src string) {
	cmd.Log(v, "*** Starting Zip decompression")
	defer cmd.Log(v, "*** Ending Zip decompression")

	zr, err := zip.OpenReader(src)
	if err != nil {
		cmd.Fatal("## " + err.Error())
	}
	defer zr.Close()

	for _, x := range zr.Reader.File {
		cmd.Log(v, "Extracting: %s", x.Name)
		c, err := x.Open()
		if err != nil {
			cmd.Fatal("## " + err.Error())
		}

		if x.FileInfo().IsDir() {
			if err = os.MkdirAll(path.Join(dst, x.Name), 0777); err != nil && !os.IsExist(err) {
				cmd.Fatal("## " + err.Error())
			}
			c.Close()
			continue
		}

		if err = os.MkdirAll(path.Join(dst, path.Dir(x.Name)), 0777); err != nil && !os.IsExist(err) {
			cmd.Fatal("## " + err.Error())
		}
		f, err := os.Create(path.Join(dst, x.Name))
		if err != nil {
			cmd.Fatal("## " + err.Error())
		}
		if _, err = io.Copy(f, c); err != nil {
			cmd.Fatal("## " + err.Error())
		}
		f.Close()
		c.Close()
	}
}
