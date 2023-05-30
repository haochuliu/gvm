package classpath

import (
	"os"
	"strings"
)

const (
	pathListSeparator = string(os.PathListSeparator)
	generalSymbol     = "*"
	ZipSuffix           = ".zip"
	JarSuffix           = ".jar"
)

type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, generalSymbol) {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(strings.ToLower(path), ZipSuffix) || strings.HasSuffix(strings.ToLower(path), JarSuffix) {
		return newCompressEntry(path)
	}
	return newDirEntry(path)
}
