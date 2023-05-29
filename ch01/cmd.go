package ch01

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	HelpFlag    bool
	VersionFlag bool
	Args        []string
	CpOption    string
	Class       string
}

func PrintUsage() {
	fmt.Printf("Usage: %s [-option] class [arg...] \n", os.Args[0])
}

func ParseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = PrintUsage
	flag.BoolVar(&cmd.HelpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.HelpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.VersionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.CpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.CpOption, "cp", "", "classpath")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.Class = args[0]
		cmd.Args = args[1:]
	}
	return cmd
}
