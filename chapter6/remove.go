package main

import (
	"log"
	"os"
)

func main() {
	err := os.RemoveAll("dir1")
	if err != nil {
		log.Fatal(err)
	}
}
