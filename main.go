package main

import (
	"fmt"
)

func main() {
	goname := New(RomanMap)
	fmt.Println(goname.FirstLast())
}
