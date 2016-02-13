package main

import "fmt"

func main(){
	for i:=0;i<5;i++{
		fmt.Println(i)
		for i:=0;i<5;i++ {
			defer fmt.Println(i)
			for i:=0;i<5;i++{
				defer fmt.Println(100-i)
			}
		}
	}
}