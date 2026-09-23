package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	cyan = "\x1b[36m"
	reset = "\x1b[0m"
)

func fs(path string, input string) {
	dirs, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, dir := range dirs {
		if strings.Contains(dir.Name(), input) {
			hl := strings.ReplaceAll(dir.Name(), input, cyan+input+reset)
			if dir.IsDir() {
				fmt.Println(hl + "/")
			} else {
				fmt.Println(hl)
			}
		}
	}
}

func run(args []string) int {
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "usage: fzfs <query>")
		return 1
	}

	fs(".", args[0])
	return 0
}

func main() {
	os.Exit(run(os.Args[1:]))
}
