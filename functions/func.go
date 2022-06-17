// package main

// import "fmt"

// func plus(a int, b int) int {
// 	return a + b
// }
// func main() {
// 	fmt.Println(plus(2, 3))
// }

// package main

// import "fmt"

// func plus()(int ,int ){
// 	return 3, 7
// }

// func main() {
// 	a, b:=plus()
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	_,x:=plus()
// 	fmt.Println(x)
// }

package main

import "fmt"

func plus(nums ...int)int{
	total:=0
	for _,num:=range nums{
		total += num
	}
	return total
}

func main(){
	sum:=plus(1, 2)
	fmt.Println(sum)
}