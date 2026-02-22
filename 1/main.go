package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a+b
}

func maymemath(a int, b int) (int, int) {
	return a+b, a*b
}

func main(){
	a := [3]int{1,2,3}
	c := &a;
	c[0] = 300

	fmt.Println(a) // [100 2 3]
	b := a
	b[0] = 100

	fmt.Println(a) // [1 2 3]
	fmt.Println(b) // [100 2 3]
		fmt.Println(a) // [1 2 3]


}