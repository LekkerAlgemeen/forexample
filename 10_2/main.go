package main

import (
	"fmt"

	"github.com/forexample/10_2/quote"
	"github.com/forexample/10_2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
