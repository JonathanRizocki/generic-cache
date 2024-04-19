package cache

// package main

// import (
// 	"fmt"
// 	"strconv"
// 	"sync"
// 	"time"

// 	"golang.org/x/sync/errgroup"
// )

// func printEverySecond(msg string) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(msg)
// 		time.Sleep(time.Second)
// 	}
// }

// func main() {
// 	// Run two goroutines
// 	for i := 0; i < 1; i++ {
// 		go printEverySecond("Hello" + strconv.Itoa(i+4))
// 		go printEverySecond("world" + strconv.Itoa(i+11))
// 	}

// 	var input string
// 	fmt.Scanln(&input)

// 	cookRecipe()
// 	cookRecipe2()
// }

// func cookRecipe2() {
// 	var g errgroup.Group
// 	g.SetLimit(2)

// 	g.Go(func() error {
// 		cookRice2()
// 		return nil
// 	})
// 	g.Go(cookCurry2)

// 	err := g.Wait()
// 	if err != nil {
// 		fmt.Printf("error %v", err)
// 	}
// }

// func cookRice2() {
// 	fmt.Println("Cooking rice2...")
// }

// func cookCurry2() error {
// 	fmt.Println("Cooking curry2...")
// 	return nil
// }

// func cookRecipe() {
// 	wg := &sync.WaitGroup{}
// 	wg.Add(2)

// 	go cookRice(wg)
// 	go cookCurry(wg)

// 	wg.Wait()
// }

// func cookRice(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println("Cooking rice...")
// 	// prep rice
// }

// func cookCurry(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println("Cooking curry...")
// 	// prep curry
// }
