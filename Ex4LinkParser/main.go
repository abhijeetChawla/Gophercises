package main

import (
	"fmt"

	linkparser "linkparserexample/linkparser"
)

func main() {
	fmt.Println(linkparser.LinkParser("ex1.html"))
	fmt.Println(linkparser.LinkParser("ex2.html"))
	fmt.Println(linkparser.LinkParser("ex3.html"))
	fmt.Println(linkparser.LinkParser("ex4.html"))
}
