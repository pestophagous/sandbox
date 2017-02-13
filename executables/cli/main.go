package main

import (
	"bufio"
	"fmt"
	"os"

	jpkg "github.com/practicum/sandbox/core/journal"
)

var (
	done    bool           = false
	journal *jpkg.Journal  = nil
	scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
)

func main() {
	processArgs()

	for !done {

		printCommandOptions()
		line, ok := getInputLine()
		if ok {
			done = !dispatch(line, journal)
		}
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
}

// processArgs must either panic or set package-global 'done' to false if any
// fatal issue is encountered.
func processArgs() {
	if len(os.Args) != 2 {
		panic("Provide exactly one argument (the path to the txt/json file.)")
	} else {
		openJournal(os.Args[1])
	}
}

// openJournal must either panic or set package-global 'done' to false if any
// fatal issue is encountered.
func openJournal(location string) {
}

func printCommandOptions() {
	/*
		TODO maybe use os.Stdin.Stat to get FileInfo,
		check the Mode() bit & os.ModeCharDevice to see if reading from pipe
	*/
	helpText := "Press one of: a (list All), d (list Done), t (list Todo)\n"
	helpText += "              n ____ (new), f __ (finish), e (exit)\n"
	helpText += "> "
	fmt.Print(helpText)
}

func getInputLine() (string, bool) {
	if !scanner.Scan() {
		done = true
		return "", false
	}

	return scanner.Text(), true
}
