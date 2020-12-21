package main

import (
	"html/template"
	"log"
	"net/http"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("view.html")
	check(err)
	err = html.Execute(writer, nil)
}


func main(){
	// net.http is the web server
	http.HandleFunc("/guestbook", viewHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
