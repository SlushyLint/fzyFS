package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	cyan = "\x1b[46m"
	reset = "\x1b[0m"
)

type pathQuery struct {
	path      string
	query     string
	recursive bool
	exact     bool
	color     bool
}

func parseArgs(args []string) (pathQuery, error) {
	result := pathQuery{recursive: false, color: true}
	positional := make([]string, 0, 2)

	for _, arg := range args {
		if strings.HasPrefix(arg, "-") && arg != "-" {
			for _, ch := range arg[1:] {
				switch ch {
				case 'r':
					result.recursive = true
				case 'x':
					result.exact = true
				case 'c':
					result.color = false
				default:
					return result, fmt.Errorf("unknown flag: -%c", ch)
				}
			}
			continue
		}
		positional = append(positional, arg)
	}

	if len(positional) != 2 {
		return result, fmt.Errorf("usage: fzx [-r] [-x] [-c] <path> <query> (flags may appear before or after positional args)")
	}

	result.path = positional[0]
	result.query = positional[1]
	return result, nil
}

func fs(path string, input string, rec bool, exact bool, color bool) {
	dirs, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, dir := range dirs {
		name := dir.Name()
		match := false
		if exact {
			match = name == input
		} else {
			match = strings.Contains(name, input)
		}
		if match {
			if exact {
				hl := name
				if color {
					hl = cyan + name + reset
				}
				if dir.IsDir() {
					fmt.Println(filepath.Join(path, hl) + "/")
				} else {
					fmt.Println(filepath.Join(path, hl))
				}
			} else {
				hl := name
				if color {
					hl = strings.ReplaceAll(name, input, cyan+input+reset)
				}
				if dir.IsDir() {
					fmt.Println(filepath.Join(path, hl) + "/")
				} else {
					fmt.Println(filepath.Join(path, hl))
				}
			}
		}
		if rec && dir.IsDir() {
			fs(filepath.Join(path, name), input, true, exact, color)
		}
	}
}

func run(args []string) int {
	parsed, err := parseArgs(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, "usage: fzx [-r] [-x] [-c] <path> <query> (flags may appear before or after positional args)")
		return 1
	}

	fs(parsed.path, parsed.query, parsed.recursive, parsed.exact, parsed.color)
	return 0
}

func main() {
	os.Exit(run(os.Args[1:]))
}
