package main

import (
	"fmt"
	"gvm/ch02"
	"gvm/ch02/classpath"
	"strings"
)

func main() {
	cmd := ch02.ParseCmd()
	if cmd.VersionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.HelpFlag || cmd.Class == "" {
		ch02.PrintUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *ch02.Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.CpOption)
	fmt.Printf("classpath: %s class:%s args:%v\n", cmd.CpOption, cmd.Class, cmd.Args)
	className := strings.Replace(cmd.Class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Cloud not find or load class %s\n", cmd.Class)
		return
	}
	fmt.Printf("Class data:%v\n", classData)
}
