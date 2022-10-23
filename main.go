package main

import (
	"fmt"
	"os"
	"log"
	"strconv"
)

func main() {
	fmt.Println("hello world")
	var args = os.Args
	//rang 截取了 左值起并且小于右值(没有右值则默认最长，没有左值默认0)的所有值组成一个新数组，idx 则是从0起 新数组长度为止
	for idx , arg := range args[1:3] {
		fmt.Println( strconv.Itoa(idx) + ":" + arg )
	}

	for idx := 0; idx < len(args); idx ++{
		log.Println(args[idx])
	} 
} 