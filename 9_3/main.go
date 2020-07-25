package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	gs := 100
	count := 0

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := count
			runtime.Gosched()
			v++
			count = v
			wg.Done()
		}()
		fmt.Println(i, count)
	}

	wg.Wait()
}
