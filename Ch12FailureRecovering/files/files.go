package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func reportPanic(){
	p := recover()
	if p != nil{
		err, ok := p.(error)
		if ok {
			fmt.Println("Error: ", err.Error())
		} else {
			panic(p)
		}
	}
}


func count (start int, end int){
	fmt.Println(start)
	if start < end {
		count(start+1, end)
	}
}

func readCatalog(path string){
	fmt.Println(path)
	names, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range names{
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			readCatalog(filePath)
		} else {
			fmt.Println(filePath)
		}
	}
}


func main(){
	defer reportPanic()
	readCatalog("my_directory")
}