package main

import (
	"fmt"
)

func main() {
	s := ", x"
	// coment with a mispeling in it on purpose.
	fmt.Println("auditor_main", s)

	{
		e, s := "e", "shadow"
		fmt.Println(s)
		fmt.Printf("%s %d\n", e, 4)
	}
}

// this funcion is badly doc'ed and not used
func unreached() {
	if true {
	} else {
		fmt.Println("unreached")
	}
}
