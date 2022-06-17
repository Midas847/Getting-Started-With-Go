package main

import "fmt"

func main() {
	b:= [5]int {1, 2, 3, 4, 5}
	for i := 0; i < 5; i++ {
		fmt.Println(b[i])
	}
	var a [2][3]int
	for i:= 0; i < 2; i++ {
		for j:= 0; j < 3; j++ {
			a[i][j] = i + j;
		}
	}
	fmt.Println("2d: ",a)
}