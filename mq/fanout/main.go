package main

import (
	"demo/services/mq"
	"fmt"
)

func main() {
	mq.ConsumerEx("fyouku.demo.fanout", "fanout", "", callback)
}

func callback(s string) {
	fmt.Printf("msg is :%s\n", s)
}
