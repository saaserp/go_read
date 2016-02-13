package main
import "fmt"
func main(){
	a := [] int {1,2,3}
	fmt.Println(a)
	modify(a)
	fmt.Println(a)



}
func modify(data [] int){
	data[0]=0
}
