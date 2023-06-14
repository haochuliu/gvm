package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClassPath Entry
}

func (this *Classpath) parseBootAndExtClasspath(option string) {
	jreDir := getJreDir(option)
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	this.bootClasspath = newWildcardEntry(jreLibPath)
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	this.extClasspath = newWildcardEntry(jreExtPath)
}

func getJreDir(option string) string {
	if option != "" && exist(option) {
		return option
	}
	if exist("./jar") {
		return "./jar"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder!")
}

func exist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (this *Classpath) parseUserClasspath(option string) {
	if option == "" {
		option = "."
	}
	this.userClassPath = newEntry(option)
}

func (this *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := this.bootClasspath.readClass(className); err == nil {
		return data, entry, nil
	}
	if data, entry, err := this.extClasspath.readClass(className); err == nil {
		return data, entry, nil
	}
	return this.userClassPath.readClass(className)
}

func (this *Classpath) String() string {
	return this.userClassPath.String()
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}
