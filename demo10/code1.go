package main

import "fmt"

func main(){
	sum:=0.0 
	var avg float64 
	xs:=[]float64{1,2,3,4,5,6}
	switch len(xs){
		case 0:avg=0
		default:
			for _,v:=range xs{
			 sum+=v
		}
		avg=sum/float64(len(xs))
	}
	fmt.Println(avg)
}