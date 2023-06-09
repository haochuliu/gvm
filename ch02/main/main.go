package main

import (
	"fmt"
	"gvm/ch02"
	"gvm/ch02/classpath"
	"strings"
)

// 测试参数  -Xjre "D:\Coding\Java\jdk-8u291\jre" java.lang.Object
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
	fmt.Printf("classpath:%v class:%v args:%v\n", cp, cmd.Class, cmd.Args)
	className := strings.Replace(cmd.Class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Cloud not find or load class %s\n", cmd.Class)
		return
	}
	fmt.Printf("class data:%v\n", classData)
}
