# Golang-json-db
Simple json based db, if you call db xd.
#Basics
```go
//First of all you have to config main json file then config backup json file. Backup is important you will see.

func main(){
 db := Db{
     path: "main.json",
     backupPath: "backup.json",
     }
 //After config we can start !!
 //We have some methods rn like write, get, backup
 //We will look at them
 db.Write("key" , "value").Backup()
 //We wrote key and value here then we backupped our datas in our backup file. Dont forget to use backup.
 db.Get("key") //will bring the data 
}



```

