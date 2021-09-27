# GoBase - Golang JSON Database
Simple Golang JSON Database module.

# Basics
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

![GoBase](https://i.pinimg.com/736x/17/34/a1/1734a186f94861a66e11f4d3c1677ee8.jpg)

