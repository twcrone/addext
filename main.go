package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	files, err := ioutil.ReadDir("./")
	check(err)
	for _, f := range files {
		fmt.Println(f.Name())
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
