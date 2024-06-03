package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	a := flag.Bool("a", false, "show hidden files")
	flag.Parse()

	for _, f := range listFiles("testdata", *a) {
		fmt.Println(f)
	}
}

func listFiles(dirname string, a bool) []string {
	var dirs []string

	files, err := os.ReadDir(dirname)

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if !a && strings.HasPrefix(f.Name(), ".") {
			continue
		}

		dirs = append(dirs, f.Name())
	}

	return dirs
}
