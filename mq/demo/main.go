package main

import (
	"demo/services/mq"
	"fmt"
)

// 简短模式及worker模式的消费者
func main() {
	mq.Consumer("", "fyouku_demo", callback)
}

func callback(s string) {
	fmt.Printf("msg is :%s\n", s)
}
