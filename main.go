package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	cmd := Command{
		Action:   "new",
		FilePath: "./",
	}
	fmt.Printf("Echo command line tool. Version: %s\n", Version)
	flag.String("h", "bb", "haha")
	flag.StringVar(&cmd.Action, "a", "new", "Command actions. Such as 'new'")
	flag.StringVar(&cmd.FilePath, "f", "./", "Create file path.")
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
