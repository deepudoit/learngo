package main

import "fmt"

//Outside main use var
//Keep least scope definition
//Types...
type coffy int
var c coffy

func main() {
	c = 12
	a := 23
	fmt.Printf("%T, %v\n", c, c)
	fmt.Printf("%d\n", a)
	a = int(c)
	fmt.Printf("%d\n", a)
}
