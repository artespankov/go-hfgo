package main

import (
	"log"
	"net/http"
)



func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message + " web!"))
	if err != nil {
		log.Fatal(err)
	}
}


func englishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello")
}

func hindiHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Namaste")
}

func frenchHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Salut")
}


func main(){
	// net.http is the web server
	http.HandleFunc("/en", englishHandler)
	http.HandleFunc("/in", hindiHandler)
	http.HandleFunc("/fr", frenchHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}