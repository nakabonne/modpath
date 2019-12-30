package main

import (
	"flag"
	"fmt"

	"github.com/nakabonne/modpath"
)

func main() {
	flag.Parse()
	dir := flag.Arg(0)
	modPath, err := modpath.Run(dir)
	if err != nil {
		fmt.Println("failed to find the module path: ", err)
		return
	}
	fmt.Println(modPath)
}
