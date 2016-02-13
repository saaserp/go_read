package main

import "fmt"

func main(){
	// for i:=0 ;i<5 ;i++{
	// 	defer fmt.Printf("%d ",i)
	// }

	for i:=0;i<5 ;i++{
		defer func(i int){
			fmt.Printf("%d ",i)
		}
	}

}