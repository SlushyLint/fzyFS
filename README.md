# fzyFS

Small experimental filesystem search tools written in Go.

This repository contains a few tiny utilities for finding files and directories by name in a shell-friendly way. They are intentionally lightweight, easy to inspect, and useful for quick local filesystem exploration.

## Status

- Experimental
- Minimal feature set
- Intended for daily use in a terminal, not as a full shell replacement

## Tools

### fz

The original non-recursive search tool.

Usage:

```bash
cd /workspaces/fzyFS/fz
go build -o fz .
./fz <path> <query>
```

Example:

```bash
./fz /path/to/project pdf
```

Behavior:

- Searches for a substring match in names
- Prints matching files and directories
- Highlights matches in cyan
- Default mode is not recursive

### fzr

A recursive version of the same idea.

Usage:

```bash
cd /workspaces/fzyFS/fzr
go build -o fzr .
./fzr <path> <query>
```

Example:

```bash
./fzr /path/to/project go
```

Behavior:

- Searches recursively through child directories
- Matches names by substring
- Highlights matches in cyan

### fzx

A recursive exact-match and flag-aware variant.

Usage:

```bash
cd /workspaces/fzyFS/fzx
go build -o fzx .
./fzx <path> <query>
./fzx -r <path> <query>
./fzx -x <path> <name>
./fzx -rx <path> <name>
./fzx -c <path> <query>
```

Examples:

```bash
./fzx /workspaces/fzyFS README.md
./fzx -r /workspaces/fzyFS README.md
./fzx -x /workspaces/fzyFS README.md
./fzx -rcx /workspaces/fzyFS README.md
```

Behavior:

- Default search is non-recursive unless `-r` is present
- `-x` searches for an exact name match instead of a substring match
- `-c` disables color output for piping and scripting
- Flags can appear before or after the positional arguments
- Matches are highlighted in cyan when color is enabled

## Project layout

```text
README.md
fz/
  main.go
  go.mod
fzr/
  main.go
  go.mod
fzx/
  main.go
  main_test.go
  go.mod
```

## Notes

- This project is intentionally small and easy to read.
- The tools are designed for fast local filesystem lookup rather than full-featured shell integration.
- The output is deliberately plain and lightweight so it can be used in pipelines and scripts.

## Roadmap

Possible follow-ups include:

- case-insensitive matching
- smarter fuzzy matching
- more flexible output formatting
- shell integration helpers

If you use these tools in their current form, think of them as a tiny experimental toolkit rather than a polished filesystem navigator.
