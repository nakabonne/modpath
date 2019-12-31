package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nakabonne/modpath"
)

func main() {
	os.Exit(main1())
}

func main1() int {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "usage: modpath [directories...]")
		flag.PrintDefaults()
	}
	flag.Parse()

	dirs := flag.Args()
	if len(dirs) == 0 {
		dirs = append(dirs, ".")
	}

	paths := make([]string, 0, len(dirs))
	for _, dir := range dirs {
		path, err := modpath.Run(dir)
		if err != nil {
			fmt.Fprintln(os.Stderr, "failed to find the module path:", err)
			return 1
		}
		paths = append(paths, path)
	}
	for _, p := range paths {
		fmt.Println(p)
	}
	return 0
}
