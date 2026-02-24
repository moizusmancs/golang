package main

import "fmt"

type Counter struct {
    Val int
}

func (c Counter) Increment() Counter {
	fmt.Println(c.Val)
	c.Val++ 
	fmt.Println(c.Val)
	return c
}
func (c Counter) PrintMe() { fmt.Println(c.Val) }


type Incrementor interface {
    Increment() Counter
	PrintMe()
}

func main() {
    c := Counter{Val: 0}
    
    var i Incrementor = c 
    
    c.Val = 100 
    i = i.(Counter).Increment() // 0,1

	fmt.Println(c.Val) // 100
	i.PrintMe() // 1
}