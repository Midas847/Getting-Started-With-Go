package main

import "fmt"

func zeroVal(fpval int){
	fpval = 0
}

func zeroptr(fptr *int){
	*fptr = 0
}

func main() {
	val := 1
	fmt.Println(val)
	zeroVal(val)
	fmt.Println("Pass by value: ",val)
	zeroptr(&val)
	fmt.Println("Pass by reference: ",val)
	fmt.Println("Address: ",&val)
}