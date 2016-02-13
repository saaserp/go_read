package main

import (
	"fmt"
	)
func main(){
	//1
	var a [3] byte
	fmt.Println(a)
	//2
	var b [2] int =[2]int{1,2}
	fmt.Println(b)
	//3
	c:=[2]int{1}
	fmt.Println(c)
	//4
	var d[2] int =[...]int{2,3}
	fmt.Println(d)
	//5
	e:=[...]int{1,2,3,4,5}
	fmt.Println(e)
}