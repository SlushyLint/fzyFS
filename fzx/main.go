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
	suffix    string
	recursive bool
	exact     bool
	color     bool
}

func parseArgs(args []string) (pathQuery, error) {
	result := pathQuery{recursive: false, color: true}
	positional := make([]string, 0, 3)

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

	if len(positional) != 2 && len(positional) != 3 {
		return result, fmt.Errorf("usage: fzx [-r] [-x] [-c] <path> <query> [suffix] (flags may appear before or after positional args)")
	}

	result.path = positional[0]
	result.query = positional[1]
	if len(positional) == 3 {
		result.suffix = positional[2]
	}
	return result, nil
}

func matches(name string, input string, suffix string, exact bool) bool {
	if suffix != "" {
		if !strings.HasSuffix(name, suffix) {
			return false
		}
		if exact {
			return name == input
		}
		return strings.Contains(name, input)
	}
	if exact {
		return name == input
	}
	return strings.Contains(name, input)
}

func highlight(name string, input string, suffix string, color bool) string {
	if !color {
		return name
	}
	text := name
	for _, token := range []string{input, suffix} {
		if token == "" {
			continue
		}
		text = strings.ReplaceAll(text, token, cyan+token+reset)
	}
	return text
}

func fs(path string, input string, suffix string, rec bool, exact bool, color bool) {
	dirs, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, dir := range dirs {
		name := dir.Name()
		match := matches(name, input, suffix, exact)
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
				hl := highlight(name, input, suffix, color)
				if dir.IsDir() {
					fmt.Println(filepath.Join(path, hl) + "/")
				} else {
					fmt.Println(filepath.Join(path, hl))
				}
			}
		}
		if rec && dir.IsDir() {
			fs(filepath.Join(path, name), input, suffix, true, exact, color)
		}
	}
}

func run(args []string) int {
	parsed, err := parseArgs(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, "usage: fzx [-r] [-x] [-c] <path> <query> [suffix] (flags may appear before or after positional args)")
		return 1
	}

	fs(parsed.path, parsed.query, parsed.suffix, parsed.recursive, parsed.exact, parsed.color)
	return 0
}

func main() {
	os.Exit(run(os.Args[1:]))
}
