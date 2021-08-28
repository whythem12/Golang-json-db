package db

import (
	"encoding/json"
	"fmt"
	"os"
	
)

func main() {
	
	db := Db{
		path: "asdf.json",
		backupPath: "backup.json",
	}
	db.Write("qwe" , "flgg")	
}
type Db struct{
	path string
	inputs map[string]interface{}
	backupPath string
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



func (db *Db) Write(key string, value interface{}) *Db {
    if db.inputs == nil {
        db.inputs = map[string]interface{}{}
    }

    db.inputs[key] = value

    data, err := json.Marshal(db.inputs)

    if err != nil {
        fmt.Println(err)
    }

    err = os.WriteFile(db.path, data, 666)

    if err != nil {
        fmt.Println(err)
    }

    return db
}
func (db *Db) Get(key string) interface{} {
	a := db.inputs[key]
	fmt.Println(a)
	return db.inputs[key]
}
func (db *Db) Backup(){
	fmt.Println(db.inputs)
	if db.inputs == nil {
        db.inputs = map[string]interface{}{}
    }
    data, err := json.Marshal(db.inputs)

    if err != nil {
        fmt.Println(err)
    }

    err = os.WriteFile(db.backupPath, data, 666)

    if err != nil {
        fmt.Println(err)
    }
}
func (db *Db) GetAll() interface{} {
	a := db.inputs
	fmt.Println(a)
	return db.inputs
}
