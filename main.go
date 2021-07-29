package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args
	var ext string
	if len(args) == 2 {
		ext = args[1]
		fmt.Println("Adding extension [" + ext + "]")
	} else {
		log.Fatal("Usage: addext [extension]")
	}
	files, err := ioutil.ReadDir("./")
	check(err)
	for _, f := range files {
		currentFilename := f.Name()
		if strings.Contains(currentFilename, ".") == false {
			newFilename := currentFilename + "." + ext
			fmt.Printf("Renaming %q to %q\n", currentFilename, newFilename)
			err := os.Rename(currentFilename, newFilename)
			check(err)
		} else {
			fmt.Printf("Skipping %q\n", currentFilename)
		}
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
