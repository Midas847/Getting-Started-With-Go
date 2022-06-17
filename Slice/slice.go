package main

import "fmt"

func main() {
	a:= make([]string, 3)
	a[0]="a"
	a[1]="b"
	a[2]="c"
	fmt.Println(a)
	c:= make([]string, len(a))
	copy(c, a)
	c = append(c, "d")
	c = append(c, "e")
	fmt.Println(c)
	twoD:= make([][]int, 3)
	for i:=0;i<3;i++ {
		leng := i + 1;
		twoD[i]=make([]int,leng)
		for j:=0;j<leng;j++ {
			twoD[i][j] = i + j
		}
	} 
	fmt.Println(twoD)
}