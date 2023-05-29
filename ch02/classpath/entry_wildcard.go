package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

func newWildcardEntry(path string) CompositeEntry {
	// 去除 * 号
	baseDir := path[:len(path) - 1]
	compositeEntry := CompositeEntry{}
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(strings.ToLower(path), JarSuffix) {
			jarEntry := newCompressEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}
	if err := filepath.Walk(baseDir, walkFn); err != nil {
		return nil
	}
	return compositeEntry
}


