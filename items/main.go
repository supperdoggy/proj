package main

import (
	"github.com/gin-gonic/gin"
	handlers2 "github.com/supperdoggy/score/items/handlers"
	db3 "github.com/supperdoggy/score/sctructs/db"
	"log"
)

func main() {
	// initializing db
	db, err := db3.Init()
	if err != nil {
		panic(err.Error())
	}
	log.Println("successfully connected to db")
	defer db.Close()

	handlers := handlers2.Handlers{DB: db}

	r := gin.Default()

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/helloworrld", handlers.HelloWorld)
	}


	if err := r.Run(":1212"); err!= nil {
		log.Println("!!! r.Run() ERROR !!!")
	}

}
