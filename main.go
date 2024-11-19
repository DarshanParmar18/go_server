package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/form", handleForm)
	http.HandleFunc("/hello", handleHello)

	fmt.Println("Starting the Server on port :8080")

	if err := http.ListenAndServe(":8080",nil); err != nil{
		log.Fatal(err)
	}
}

func handleForm( w http.ResponseWriter, req *http.Request)  {

	if err := req.ParseForm(); err != nil	{
		fmt.Fprintf(w, "parseForm() err : %+v\n", err)
	}
	
	fmt.Fprintln(w, "Post request successfull!");
	name := req.FormValue("name")
	address := req.FormValue("address")
	
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func handleHello( w http.ResponseWriter, req *http.Request)  {

	if req.URL.Path != "/hello" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	if req.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintln(w, "hello")
}