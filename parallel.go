package main

import (
	"fmt"
	"strconv"
	"sync"
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
		go printEverySecond("world" + strconv.Itoa(i+11))
	}

	var input string
	fmt.Scanln(&input)

	cookRecipe()
}

func cookRecipe() {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go cookRice(wg)
	go cookCurry(wg)

	wg.Wait()
}

func cookRice(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Cooking rice...")
	// prep rice
}

func cookCurry(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Cooking curry...")
	// prep curry
}
