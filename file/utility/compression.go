package utility

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	cmd "cookbook"
	"cookbook/file/forensics"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// CompressNew compresses the files at src into dst based on method.
func CompressNew(v bool, dst string, src ...string) {
	cmd.Log(v, "\n*** Starting compression\n")
	defer cmd.Log(v, "Ending compression ***\n")
	if path.IsAbs(dst) {
		log.Fatalf("Destination %s must be relative", dst)
	}
	n := strings.Split(path.Base(dst), ".")
	if len(n) < 2 {
		log.Fatalf("Destination %s cannot be a directory\n", dst)
	}

	if err := os.MkdirAll(path.Dir(path.Clean(dst)), 0777); err != nil && !os.IsExist(err) {
		log.Fatal(err.Error())
	}
	cmd.Log(v, "- Creating temp dir\n")
	tmp, err := os.MkdirTemp(path.Dir(dst), n[0])
	if err != nil {
		log.Fatal(err.Error())
	}
	defer os.RemoveAll(tmp)

	cmd.Log(v, "- Creating the destination file at: %s\n", dst)
	f, err := os.Create(dst)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	switch filepath.Ext(path.Base(dst)) {
	case ".tar":
		compressTar(v, tmp, f, src)
	case ".gz":
		compressGZ(v, tmp, f, src)
	case ".zip":
		compressZip(v, tmp, f, src)
	default:
		cmd.Log(v, "* Unknown compression method: %v\n", n[1:])
	}
}

func compressTar(v bool, tmp string, w io.Writer, src []string) {
	cmd.Log(v, "\n*** Starting Tar compression\n")
	defer cmd.Log(v, "Ending Tar compression ***\n")

	for _, x := range src {
		forensics.ExtractCopy(v, path.Join(tmp, path.Base(x)), x)
	}
	a := forensics.Enumerate(v, "", false, tmp)

	t := tar.NewWriter(w)
	defer t.Close()
	for _, x := range a.GetPaths(0) {
		buf := bytes.NewBuffer(nil)
		forensics.Extractor(v, buf, x)
		t.Write(buf.Bytes())
	}
}

func compressGZ(v bool, tmp string, w io.Writer, src []string) {
	cmd.Log(v, "\n*** Starting GZip compression\n")
	defer cmd.Log(v, "Ending GZip compression ***\n")

	for _, x := range src {
		n := path.Clean(strings.Join(strings.Split(x, ".."), ""))
		forensics.ExtractCopy(v, path.Join(tmp, n))
	}
	a := forensics.Enumerate(v, "", false, tmp)

	g := gzip.NewWriter(w)
	defer g.Close()
	for _, x := range a.GetPaths(0) {
		buf := bytes.NewBuffer(nil)
		forensics.Extractor(v, buf, x)
		g.Write(buf.Bytes())
	}
}

func compressZip(v bool, tmp string, w io.Writer, src []string) {
	cmd.Log(v, "\n*** Starting Zip compression\n")
	defer cmd.Log(v, "Ending Zip compression ***\n")

	for _, x := range src {
		n := path.Clean(strings.Join(strings.Split(x, ".."), ""))
		forensics.ExtractCopy(v, path.Join(tmp, n))
	}
	a := forensics.Enumerate(v, "", false, tmp)

	z := zip.NewWriter(w)
	defer z.Close()
	for _, p := range a.GetPaths(0) {
		r, err := z.Create(p[len(tmp)-1:])
		if err != nil {
			log.Fatal(err.Error())
		}
		buf := bytes.NewBuffer(nil)
		forensics.Extractor(v, buf, p)
		r.Write(buf.Bytes())
	}
}

//func Decompress(v bool, dst, method, src string) {
//	f, err := os.Open(src)
//	if err != nil {
//		log.Fatal(err.Error())
//	}
//
//	switch method {
//	case "tar":
//		decompressTar(v, f, dst)
//	case "tar.gz":
//		// decompressGZ(f, dst)
//		decompressTar(v, f, dst)
//	case "gz":
//		fallthrough
//	case "gzip":
//		decompressGZ(v, f, dst)
//	case "zip":
//		decompressZip(v, f, dst)
//	default:
//		log.Println(v, "* Unknown method: %s", method)
//	}
//}
//
//func decompressTar(v bool, r io.Reader, dst string) {
//
//}
//
//func decompressGZ(v bool, r io.Reader, dst string) {
//
//}
//
//func decompressZip(v bool, r io.Reader, dst string) {
//
//}
