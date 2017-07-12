package main

import (
	"fmt"
	goname "github.com/folkengine/goname"
)

func main() {
	goname := goname.New(goname.RomanMap)
	fmt.Println(goname.FirstLast())
}
