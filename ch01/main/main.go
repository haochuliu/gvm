package main

import (
	"fmt"
	"gvm/ch01"
)

func main() {
	cmd := ch01.ParseCmd()
	if cmd.VersionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.HelpFlag || cmd.Class == "" {
		ch01.PrintUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *ch01.Cmd) {
	fmt.Printf("classpath: %s class:%s args:%v\n", cmd.CpOption, cmd.Class, cmd.Args)
}
