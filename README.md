[![Build Status](https://travis-ci.org/jeffsmith82/gofuzzy.svg?branch=master)](https://travis-ci.org/jeffsmith82/gofuzzy)

gofuzzy
=======

Golang Fuzzy matching Library

This is just a testing repo to learn how to use github properly.


Levenshtein distance Example
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

Soundex Example
-------

	package main
	
	import (
		"fmt"
		"github.com/jeffsmith82/gofuzzy"
	)

	func main(){
		fmt.Println(gofuzzy.Soundex("Robert"))
		fmt.Println(gofuzzy.Soundex("Rupert"))
		fmt.Println(gofuzzy.Soundex("Rubin"))
		fmt.Println(gofuzzy.Soundex("ashcroft"))
		fmt.Println(gofuzzy.Soundex("ashcraft"))
		fmt.Println(gofuzzy.Soundex("moses"))
		fmt.Println(gofuzzy.Soundex("O'Mally"))
		fmt.Println(gofuzzy.Soundex("d jay"))
		fmt.Println(gofuzzy.Soundex("R2-D2"))
		fmt.Println(gofuzzy.Soundex("12p2"))
		fmt.Println(gofuzzy.Soundex("naïve"))
		fmt.Println(gofuzzy.Soundex(""))
		fmt.Println(gofuzzy.Soundex("bump\t"))
	}
