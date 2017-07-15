package main

import (
	"flag"
	"fmt"
	goname "github.com/folkengine/goname"
)

func main() {

	elvenFlag := flag.Bool("elven", false, "Create Elven names")
	fantasyFlag := flag.Bool("fantasy", false, "Create Fantasy names")
	goblinFlag := flag.Bool("goblin", false, "Create Goblin names")
	romanFlag := flag.Bool("roman", true, "Create Roman names")
	flag.Parse()

	var myname goname.Goname

	if *elvenFlag {
		myname = goname.New(goname.ElvenMap)
	} else if *fantasyFlag {
		myname = goname.New(goname.FantasyMap)
	} else if *goblinFlag {
		myname = goname.New(goname.GoblinMap)
	} else if *romanFlag {
		myname = goname.New(goname.RomanMap)
	}

	fmt.Println(myname.FirstLast())
}
