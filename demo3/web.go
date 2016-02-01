package main

import (
    "fmt"
	"html/template"
	"net/http"
)
func main(){
	http.HandleFunc("/",index)
	http.ListenAndServe("localhost:80",nil)

}
func index(rw http.ResponseWriter,req *http.Request){
	message :=""
	if req.Method == "POST" {
		req.ParseForm()
		if checkUsername(req.Form["username"][0]) {
			fmt.Fprint(rw,"username:",req.Form["username"][0],"password:",req.Form["password"][0])
		}else{
			message = "username is not valide"
		}
		t, _ := template.ParseFiles("index.tpl")
		t.Execute(rw,message)
	}


}
func checkUsername(username string) bool{
	if len(username)>=6 && len(username)<=16{
		return true
	}
	return false;
}
