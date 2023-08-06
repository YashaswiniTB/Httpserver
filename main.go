package main

import (
	"fmt"
	"log"
	"net/http"
	"time"


)
func hellohandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path !="/hello"{
http.Error(w,"404 status not found",http.StatusNotFound)
return
	}
	if r.Method!="GET"{
		http.Error(w,"method not supported",http.StatusNotFound)
		return
	}
	fmt.Println(w,"hello")
}
func formhandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err!= nil{
		fmt.Fprintf(w, "parseform() err: %v",err)
		return
	}
     fmt.Fprintf(w, "POST REQUEST IE SUCCESSFUL")
	 name :=r.FormValue("name")
	 address:=r.FormValue("address")
	 fmt.Fprintf(w, "name =%s\n",name)
	 fmt.Fprintf(w, "name =%s\n",address)
}
func main(){
fileserver := http.FileServer(http.Dir("./static"))
http.Handle("/",fileserver)
http.HandleFunc("/forms",formhandler)
http.HandleFunc("/hello",hellohandler)
srv := &http.Server{
	ReadTimeout: 3*time.Second,
	WriteTimeout: 6*time.Second,
	Addr: "localhost:8080",
	IdleTimeout: 15*time.Second,
	ReadHeaderTimeout: 3*time.Second,

}
fmt.Printf("starting server at port 8080")
if err :=srv.ListenAndServe(); err!=nil{
	log.Fatal(err)
}


}
