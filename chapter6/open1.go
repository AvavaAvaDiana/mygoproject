package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("open.txt")
	if err != nil {
		fmt.Printf("打开文件错误:%v\n", err)
	}

	fmt.Println(file)

	err = file.Close()
	if err != nil {
		fmt.Printf("关闭文件错误: %v\n", err)
	}
}
