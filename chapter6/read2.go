package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	filePath := "login.gtpl"
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("读取文件出错:%v", err)
	}
	fmt.Printf("%v\n", content)
	fmt.Printf("%v\n", string(content))
}
