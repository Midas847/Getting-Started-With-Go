package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "hello world"
	fmt.Println(len(s))
	fmt.Println("Rune count:", utf8.RuneCountInString(s))
	for i,value := range s {
		fmt.Printf("%#U starts at %d\n",value,i)
	}
}