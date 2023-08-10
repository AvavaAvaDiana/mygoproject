package main

import (
	"fmt"
	"os"
)

func main() {
	fp, err := os.Create("./chmod1.txt")
	if err != nil {
		fmt.Println("文件创建失败")
		return
	}
	defer fp.Close()

	fileInfo, err := os.Stat("./chmod1.txt")
	if err != nil {
		fmt.Println("获取文件信息失败")
		return
	}
	fileMode := fileInfo.Mode()
	fmt.Println("初始文件权限:", fileMode)

	os.Chmod("./chmod1.txt", 0777)

	fileInfo, err = os.Stat("./chmod1.txt")
	if err != nil {
		fmt.Println("获取文件信息失败")
		return
	}
	fileMode = fileInfo.Mode()
	fmt.Println("修改后文件权限:", fileMode)
}
