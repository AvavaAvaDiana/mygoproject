package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("WriteString.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.WriteString("Go Web 编程实战从入门到精通")
}
