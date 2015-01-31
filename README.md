gofuzzy
=======

Example
-------

	package main
	
	import (
		"fmt"
		"github.com/jeffsmith82/gofuzzy"
	)
	
	func main(){
		fmt.Println(gofuzzy.Ld("Kitten","sitting"))
		fmt.Println(gofuzzy.Ld("Test","testing"))
		fmt.Println(gofuzzy.Ld("Miles","Apart"))
	}


Golang Fuzzy matching Library
