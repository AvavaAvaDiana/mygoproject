package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strconv"
)

type ArgsTwo struct {
	X, Y int
}

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8808")
	if err != nil {
		log.Fatal("在这里地方发生了错误：DialHTTP", err)
	}

	i1, _ := strconv.Atoi(os.Args[1])
	i2, _ := strconv.Atoi(os.Args[2])
	args := ArgsTwo{i1, i2}
	var reply int
	err = client.Call("Algorithm.Sum", args, &reply)
	if err != nil {
		log.Fatal("Call Sum algorithm error:", err)
	}
	fmt.Printf("Algorithm和为:%d+%d=%d\n", args.X, args.Y, reply)
}
