package main

import "fmt"

func fact(n int)int{
	if(n == 0){
		return 1
	}
	return (n * fact(n-1))
}
func main() {
	f := fact(5)
	fmt.Println(f)

	var fun func(n int)int
	fun = func(n int)int{
		if(n < 2){
			return n
		}
		return fun(n-1) + fun(n-2)
	}
	fmt.Println(fun(7))
}