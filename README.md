# fzyFS

Beta / experimental filesystem search tools written in Go.

This repository contains a pair of lightweight utilities meant to help find files and directories by fuzzy matching names. The project is intentionally small, rough around the edges, and designed for experimentation rather than production use.

## Status

- Beta
- Early-stage
- Minimal feature set
- Subject to change without notice

## Tools

### fz

A simple fuzzy filename search utility. It walks a directory tree and highlights matches in the terminal.

Usage:

```bash
cd go/fz
go build -o fz .
./fz <path> <query>
./fz -r <path> <query>
```

Examples:

```bash
./fz . pdf
./fz -r /path/to/project go
```

Behavior:

- Searches for a substring match in file or directory names
- Prints matching entries
- `-r` enables recursive search
- Matching text is highlighted in cyan

### fzfs

A simpler version that searches the current directory only.

Usage:

```bash
cd go/fzfs
go build -o fzfs .
./fzfs <query>
```

Example:

```bash
./fzfs notes
```

## Project layout

```text
README.md
fz
go/
  fz/
    main.go
    go.mod
  fzfs/
    main.go
    go.mod
```

## Notes

- This is a beta project and not a polished shell replacement.
- The current implementation is intentionally small and easy to inspect.
- The goal is a fast, minimal search experience for local filesystem browsing.

## Roadmap

Planned ideas for future iterations include:

- better fuzzy matching behavior
- more shell-like output styling
- optional case-insensitive search
- improved recursion and filtering
- a more polished CLI interface

If you're using this project in its current form, treat it as an experimental utility rather than a production-ready filesystem tool.
