package qparser

import "strings"

type Result struct {
	Namespace string
	Query     string
}

func Parse(args []string) Result {
	if len(args) == 0 {
		return Result{Namespace: ""}
	}

	if len(args) == 1 {
		return Result{Namespace: args[0]}
	}

	return Result{Namespace: args[0], Query: strings.ReplaceAll(strings.Join(args[1:], " "), " ", "+")}
}

func (r Result) GetNamespace() string {
	return r.Namespace
}

func (r Result) GetQuery() string {
	return r.Query
}
