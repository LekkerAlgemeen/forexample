package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	var m sync.Mutex

	gs := 100
	count := 0

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			m.Lock()
			v := count
			v++
			count = v

			fmt.Println(count)
			m.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
}
