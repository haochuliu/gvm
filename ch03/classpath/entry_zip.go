package classpath

import (
	"archive/zip"
	"errors"
	"io"
	"path/filepath"
)

type ZipEntry struct {
	absDir string
	zipRc  *zip.ReadCloser
}

func newCompressEntry(path string) *ZipEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absDir, nil}
}

func (this *ZipEntry) readClass (className string) ([]byte, Entry, error) {
	if this.zipRc == nil {
		err := this.openJar()
		if err != nil {
			return nil ,nil, err
		}
	}
	classFile := this.findClass(className)
	if classFile == nil {
		return nil, nil, errors.New("class not found: " + className)
	}
	data, err := readClass(classFile)
	return data, this, err
}

func (this *ZipEntry) openJar() error {
	r, err := zip.OpenReader(this.absDir)
	if err == nil {
		this.zipRc = r
	}
	return err
}

func (this *ZipEntry) findClass (className string) *zip.File {
	for _, f := range this.zipRc.File {
		if f.Name == className {
			return f
		}
	}
	return nil
}

func readClass(classFile *zip.File) ([]byte, error) {
	rc, err := classFile.Open()
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(rc)
	if err = rc.Close(); err != nil {
		return nil, err
	}
	return data, nil
}

func (this *ZipEntry) String() string {
	return this.absDir
}

