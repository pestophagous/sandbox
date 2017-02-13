package main

import (
	"fmt"
	"strings"

	// "github.com/practicum/sandbox/core/item"
	jpkg "github.com/practicum/sandbox/core/journal"
)

const genericError = "Unable to parse: %s\n"

func dispatch(line string, jfile *jpkg.Journal) bool {
	trimmed := strings.TrimSpace(line)
	if len(trimmed) == 0 {
		fmt.Printf(genericError, line)
		fmt.Println()
		return true // main loop should continue
	}

	var c byte = trimmed[0]

	if c == 'e' || c == 'E' {
		return false // will EXIT main
	}

	switch {
	default: // default can appear anywhere - it's always the last resort
		fmt.Printf(genericError, line)
	case c == 'a' || c == 'A':
		listAll(line, jfile)
	case c == 'd' || c == 'D':
		listDone(line, jfile)
	case c == 't' || c == 'T':
		listTodo(line, jfile)
	case c == 'n' || c == 'N':
		newItem(line, jfile)
	case c == 'f' || c == 'F':
		itemIsFinished(line, jfile)
	}

	fmt.Println()

	return true // main loop should continue
}

func listAll(line string, jfile *jpkg.Journal) {
	fmt.Println("listAll")
}

func listDone(line string, jfile *jpkg.Journal) {
	fmt.Println("listDone")
}

func listTodo(line string, jfile *jpkg.Journal) {
	fmt.Println("listTodo")
}

func newItem(line string, jfile *jpkg.Journal) {
	fmt.Println("newItem")
}

func itemIsFinished(line string, jfile *jpkg.Journal) {
	fmt.Println("itemIsFinished")
}
