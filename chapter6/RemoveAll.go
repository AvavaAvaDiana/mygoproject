package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	os.MkdirAll("test_remove", 0777)
	err := os.RemoveAll("test1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("create dir :test_remove")

	_, err = os.Create("./test_remove/test_remove1.txt")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("creat file:test_remove1.txt")

	fmt.Println()
}
