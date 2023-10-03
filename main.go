package main

import (
	"filesize/fs"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
		fmt.Println("Usage: filesize <file>")
		return
	}

	firstArg := args[0]
	// check if file exists
	if _, err := os.Stat(firstArg); os.IsNotExist(err) {
		fmt.Println("File does not exist")
		return
	}

	bold := "\033[1m"
	reset := "\033[0m"
	blue := "\033[34m"

	fmt.Print(bold + blue + firstArg + reset)
	fmt.Print(" ")
	// get file stat
	stat, err := fs.GetStat(firstArg)
	if err != nil {
		fmt.Print("ERR\n")
		fmt.Println(err)
		return
	}
	// get file size
	size, err := stat.GetReadableSize()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(bold + size + reset)
}
