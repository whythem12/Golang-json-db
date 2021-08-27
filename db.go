package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	
	db := Db{
		path: "asdf.json",
	}

	
}


type Db struct{
	path string
	things map[string]interface{}
	
}


func (db *Db) Create(fileName Db ){

	a , err := os.Create(fileName.path)
	defer a.Close()
	if err != nil{
		panic(err)
	}
	
}
func (db *Db) WriteString(file Db , input string) {
	f , err := os.Create(file.path)
	if err != nil{
		panic(err)
	}
	f.WriteString(input)
//useless but who cares lol
}



func (db *Db) Check(path Db) (inf os.FileInfo, err error) {

	if inf, err = os.Stat(path.path); os.IsNotExist(err) {
		inf, err = os.Stat(path.path)
	}
	return
}
