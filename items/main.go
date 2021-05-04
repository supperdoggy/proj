package main

import (
	"github.com/gin-gonic/gin"
	db2 "github.com/supperdoggy/score/users/db"
	"log"
)

func main() {
	// initializing db
	db, err := db2.Init()
	if err != nil {
		panic(err.Error())
	}
	log.Println("successfully connected to db")
	defer db.Close()

	r := gin.Default()

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/helloworrld")
	}


	if err := r.Run(); err!= nil {
		log.Println("!!! r.Run() ERROR !!!")
	}

}
