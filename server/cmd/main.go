package main

import (
	"log"
	"server/db"
)


func main(){
	_,err:=db.NewDatabse()

	if err!=nil{
		log.Fatalf("could not initialize database connection %s",err)
	}
}