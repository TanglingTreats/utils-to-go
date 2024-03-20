package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	switch argv := len(args); argv {
	case 0:
		noArgs()
	case 1:
		checkUtils(args[0])
	default:
		fmt.Println("No such options. Bye bye")
	}
}

func noArgs() {
	fmt.Println("walk through utils")
}

func checkUtils(utilType string) {

	util, ok := utilMap[strings.ToLower(utilType)]
	if !ok {
		fmt.Println("util type does not exist")
		return
	}

	fmt.Printf("Jump to utils if exist: %v\n", util)
}
