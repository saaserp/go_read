package main

import "fmt"

func main(){
	str:= 'A'
	 for i:=1;i<=20;i++{
		fmt.Println(str)
		str+='A'
	}
}