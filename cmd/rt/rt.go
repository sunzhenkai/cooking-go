package main

import (
	"evaluate-go/pkg/rt"
	"time"
)

func main() {
	go rt.Task()
	go rt.Task()
	println("YES")

	time.Sleep(2 * time.Second)
}
