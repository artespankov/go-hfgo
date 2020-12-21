package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func repeat(s string){
	for i := 0; i < 25; i++ {
	fmt.Print(s)
	}
}


type Page struct {
	size int
	URL string
}

func respSize(url string, channel chan Page){
	fmt.Println("Getting", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		log.Fatal(err)
	}
	channel <- Page{size: len(body), URL: url}
}

func main(){
	pages := make(chan Page)
	urls := []string{"http://example.com", "http://golang.org", "http://golang.org/doc"}
	for _, url := range urls {
		go respSize(url, pages)
	}
	for i:=0; i < len(urls); i++ {
		fmt.Println(<-pages)
	}
}