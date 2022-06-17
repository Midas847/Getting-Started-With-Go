package main

import "fmt"

func main() {
	m:= make(map[string]int)
	m["a"]=1
	m["b"]=2
	m["c"]=3
	v:=m["a"]
	fmt.Println(v)
}