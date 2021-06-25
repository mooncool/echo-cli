package main

import (
	"fmt"
	"os"
)

func genFolder(path string) {
	info, err := os.Stat(path)
	fmt.Println(info)
	// if !info.IsDir() {
	// 	panic(err)
	// }

	if os.IsNotExist(err) {
		fmt.Printf("folder %s doesn't exist, creating...\n", path)
		if err = os.MkdirAll(path, 0755); err != nil {
			panic(err)
		}
		fmt.Printf("folder %s is created\n", path)
	}

	// if err = os.MkdirAll(path, 0755); err != nil {
	// 	panic(err)
	// }
}
