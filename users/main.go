package main

import (
	"github.com/gin-gonic/gin"
	db3 "github.com/supperdoggy/score/sctructs/db"
	handlers2 "github.com/supperdoggy/score/users/handlers"
	"log"
)

func main() {
	// initializing db
	db, err := db3.InitUsersDB()
	if err != nil {
		panic(err.Error())
	}
	log.Println("successfully connected to db")
	defer db.Close()

	// initializing handlers
	handlers := handlers2.Handlers{
		DB: db,
	}

	// initializing routes
	r := gin.Default()

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/createuser", handlers.CreateUser)
		apiv1.GET("/getallusers", handlers.GetAllUsers)
	}

	if err := r.Run(":1321"); err != nil {
		log.Println("r.Run() error!!!")
	}

}
