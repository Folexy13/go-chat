package main

import (
	"log"
	"server/db"
	"server/internals/user"
	"server/internals/ws"
	"server/router"
)

func main() {
	dbConn, err := db.NewDatabse()

	if err != nil {
		log.Fatalf("could not initialize database connection %s", err)
	}
	userRep := user.NewRespository(dbConn.GetDB())
	userSrvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSrvc)

	hub :=ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	router.InitRouter(userHandler,wsHandler)
	router.Start("0.0.0.0:8081")
}
