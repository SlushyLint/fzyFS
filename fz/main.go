package main

import (
	"fmt"
	"os"
	"strings"
    "path/filepath"
)

const (
	cyan = "\x1b[46m"
	reset = "\x1b[0m"
)

func fs(path string, input string, rec bool) {
	dirs, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, dir := range dirs {
		if strings.Contains(dir.Name(), input) {
			hl := strings.ReplaceAll(dir.Name(), input, cyan+input+reset)
			if dir.IsDir() {
                fmt.Println(filepath.Join(path, hl) + "/")
			} else {
                fmt.Println(filepath.Join(path, hl))
			}
		}
        if rec && dir.IsDir() {
            fs(path+"/"+dir.Name(), input, true)
        }
	}
}

func run(args []string) int {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "usage: fz <path> <query> | fz -r <path> <query>")
		return 1
	}

	if args[0] == "-r" {
		if len(args) != 3 {
			fmt.Fprintln(os.Stderr, "usage: fz -r <path> <query>")
			return 1
		}
		fs(args[1], args[2], true)
		return 0
	}

	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: fz <path> <query>")
		return 1
	}

	fs(args[0], args[1], false)
	return 0
}

func main() {
	os.Exit(run(os.Args[1:]))
}
