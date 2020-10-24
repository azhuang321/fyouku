package main

import (
	"demo/services/mq"
	"fmt"
)

func main() {
	mq.Consumer("", "fyouku_demo", callback)
}

func callback(s string) {
	fmt.Printf("msg is :%s\n", s)
}
