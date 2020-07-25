package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("GOARCH:", runtime.GOARCH)
	fmt.Println("GOOS:", runtime.GOOS)
}
