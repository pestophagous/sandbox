package main

import (
	"fmt"
)

func main() {
	s := ", x"
	// coment with a mispeling in it on porpose.
	fmt.Println("auditor_main", s)
	s = " ineffectual assignment"
	{
		s := "shadow"
		fmt.Println(s)
		fmt.Print("%s %d") // should be linted
	}
}

// this funcion is badly doc'ed and not used
func unreached() {
	fmt.Println("unreached")
}
