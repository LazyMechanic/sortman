package types

import (
	"fmt"
	"io"
	"strings"
)

type Action string

type Request struct {
	Patterns     []string
	Exclusions   []string
	Files        []string
	InDirectory  string
	OutDirectory string
}
func (r *Request) Fprint(writer io.Writer) {
	var patterns string
	if len(r.Patterns) > 0 {
		patterns = "\n  " + strings.Join(r.Patterns, "\n") + "\n"
	}

	var exclusions string
	if len(r.Exclusions) > 0 {
		exclusions = "\n  " + strings.Join(r.Exclusions, "\n") + "\n"
	}


	var outStrings = []string {
		fmt.Sprintf("Patterns: {%s}", patterns),
		fmt.Sprintf("Exclusions: {%s}", exclusions),
		fmt.Sprintf("InDirectory:  %q", r.InDirectory),
		fmt.Sprintf("OutDirectory: %q", r.OutDirectory),
	}

	for _, str := range outStrings {
		_, _ = writer.Write([]byte(str))
		_, _ = writer.Write([]byte("\n"))
	}
}

type Config struct {
	InDirectory  string
	OutDirectory string
	Requests     []Request
	Action       Action
}

type Cancel struct{}

func (e *Cancel) Error() string {
	return "Nothing happen"
}

type Execute struct{}

func (e *Execute) Error() string {
	return ""
}
