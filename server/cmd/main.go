package main

import (
	"log"
	"server/db"
	"server/internals/user"
	"server/router"
)


func main(){
	dbConn,err:=db.NewDatabse()

	if err!=nil{
		log.Fatalf("could not initialize database connection %s",err)
	}
	userRep:= user.NewRespository(dbConn.GetDB())
	userSrvc:=user.NewService(userRep)
	userHandler:=user.NewHandler(userSrvc)

	router.InitRouter(userHandler)
	router.Start("0.0.0.0:8080")
}