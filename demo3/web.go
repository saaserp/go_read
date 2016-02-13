package main

import (
    "fmt"
	"html/template"
	"net/http"
)
func main(){
    fmt.Println("hello word")
	http.HandleFunc("/",index)
    fmt.Println("2")
	http.ListenAndServe("localhost:80",nil)
    fmt.Println("3")

}
func index(rw http.ResponseWriter,req *http.Request){
	message :=""
    fmt.Println("4")
	if req.Method == "POST" {
    
		req.ParseForm()
		if checkUsername(req.Form["username"][0]) {
			fmt.Fprint(rw,"username:",req.Form["username"][0],"password:",req.Form["password"][0])
        return              
		}else{
			message = "username is not valide"
		}
    }
		t, _ := template.ParseFiles("index.tpl")
		t.Execute(rw,message)
}



func checkUsername(username string) bool{
    fmt.Println("check")
	if len(username)>=6 && len(username)<=16{
		return true
	}
	return false
}
