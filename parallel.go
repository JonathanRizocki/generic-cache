package main

import (
	"fmt"
	"strconv"
	"time"
)

func printEverySecond(msg string) {
	for i := 0; i < 10; i++ {
		fmt.Println(msg)
		time.Sleep(time.Second)
	}
}

func main() {
	// Run two goroutines
	for i := 0; i < 3; i++ {

		go printEverySecond("Hello" + strconv.Itoa(i+4))
		go printEverySecond(string("world" + strconv.Itoa(i+11)))
	}

	var input string
	fmt.Scanln(&input)
}
