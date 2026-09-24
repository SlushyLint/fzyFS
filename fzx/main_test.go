package main

import "testing"

func TestParseArgsFlagOrder(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want pathQuery
	}{
		{name: "prefix flags", args: []string{"-rcx", "/", "ls"}, want: pathQuery{path: "/", query: "ls", recursive: true, exact: true, color: false}},
		{name: "suffix flags", args: []string{"/", "ls", "-x"}, want: pathQuery{path: "/", query: "ls", recursive: false, exact: true, color: true}},
		{name: "default non-recursive", args: []string{"/", "ls"}, want: pathQuery{path: "/", query: "ls", recursive: false, exact: false, color: true}},
		{name: "no color flag", args: []string{"-c", "/", "ls"}, want: pathQuery{path: "/", query: "ls", recursive: false, exact: false, color: false}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseArgs(tt.args)
			if err != nil {
				t.Fatalf("parseArgs(%v) returned error: %v", tt.args, err)
			}
			if got != tt.want {
				t.Fatalf("parseArgs(%v) = %#v, want %#v", tt.args, got, tt.want)
			}
		})
	}
}

func TestParseArgsRejectsMissingArguments(t *testing.T) {
	if _, err := parseArgs([]string{"-x"}); err == nil {
		t.Fatal("expected missing positional args error")
	}
}
