package main

import "fmt"

func main() {
	arr:= make([]int, 3)
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr = append(arr,4, 5, 6)
	for _,num := range arr {
		fmt.Println(num)
	}
	m:= make(map[string]int)
	m["a"]=1
	m["b"]=2
	m["c"]=3
	for k,v := range m {
		fmt.Println("%s -> %s",k,v)
	}
}