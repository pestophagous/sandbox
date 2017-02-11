package main

import (
	"fmt"
)

func main() {
	s := ", x"
	// misspelled words not found by reportcard (Feb 2017): coment, mispeling
	fmt.Println("auditor_main", s)

	{
		e, s := "e", "shadow" // s is possible accidental shadow but not caught
		fmt.Println(s)
		fmt.Printf("%s %s %d\n", e, s, 4)
	}
}

// reportcard must not be auditing docu-comment style
// this funcion is badly doc'ed and not used
func unreached() {
	if true {
	} else {
		fmt.Println("unreached") // this is also not flagged
	}
}
