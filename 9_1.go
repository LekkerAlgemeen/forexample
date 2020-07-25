package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println(runtime.NumCPU())

	wg.Add(2)

	go print1()
	go print2()

	wg.Wait()
}

func print1() {
	fmt.Println("Hi")
	wg.Done()
}

func print2() {
	fmt.Println("Hello")
	wg.Done()
}
