package main

import "fmt"

func seq() func()int{
	i := 0
	return func()int{
		i++;
		return i
	}
}
func main() {
	a := seq()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
}