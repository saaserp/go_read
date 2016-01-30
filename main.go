package main

import "fmt"
import "net/http"
 

func main() {
	http.HandlerFunc("/",indexHandler)
	http.HandlerFunc("/logout",logout)
	http.ListenAndServer("localhost:80",nil)

}

func indexHandler(rw http.ResponseWriter,req *http.Request){
	req.ParseForm()
	if len(req.Form["name"]) > 0 {
		fmt.Fprint(rw,"hello,",req.Form("name")[0])
	}else{
		fmt.Fprint(rw,"hello,world")
	}
}

func logout(rw http.ResponseWriter,req *http.Request){
	fmt.Fprint(rw,"logout")
}