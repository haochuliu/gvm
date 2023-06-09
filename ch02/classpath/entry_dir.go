package classpath

import (
	"os"
	"path/filepath"
)

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

func (this *DirEntry) readClass (className string) ([]byte, Entry, error) {
	fileName := filepath.Join(this.absDir, className)
	data, err := os.ReadFile(fileName)
	return data, this, err
}

func (this *DirEntry) String() string {
	return this.absDir
}
