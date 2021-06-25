package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	cmd := Command{
		Action:     cmdActionInitWebProject,
		ProjectDir: "./",
	}
	fmt.Printf("Echo project command line tool. Version: %s\n", Version)

	flag.String("h", "bb", "haha")
	flag.StringVar(&(cmd.Action), "a", cmdActionInitWebProject, "Command actions. Such as 'init'")
	flag.StringVar(&(cmd.ProjectDir), "d", defaultWorkingDir, "Project directory.")
	fmt.Println("command:", &cmd)

	flag.Parse()

	if flag.NArg() > 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	for i := range flag.Args() {
		fmt.Println("//", i)
	}
	// execute the command
	executeCommand(cmd)
}
