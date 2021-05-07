package main

import (
	"github.com/gin-gonic/gin"
	"github.com/supperdoggy/score/sctructs"
	db3 "github.com/supperdoggy/score/sctructs/db"
	usersdata "github.com/supperdoggy/score/sctructs/service/users"
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

	apiv1 := r.Group(sctructs.ApiV1)
	{
		apiv1.POST(usersdata.CreatePath, handlers.CreateUser)
		apiv1.GET(usersdata.GetAllPath, handlers.GetAllUsers)
		apiv1.POST(usersdata.FindPath, handlers.Find)
		apiv1.POST(usersdata.DeletePath, handlers.Delete)
		apiv1.POST(usersdata.FindWithPasswordPath, handlers.FindWithPassword)
	}

	if err := r.Run(":12321"); err != nil {
		log.Println("r.Run() error!!!")
	}

}
