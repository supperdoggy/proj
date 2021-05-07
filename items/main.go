package main

import (
	"github.com/gin-gonic/gin"
	handlers2 "github.com/supperdoggy/score/items/handlers"
	"github.com/supperdoggy/score/sctructs"
	db3 "github.com/supperdoggy/score/sctructs/db"
	itemsdata "github.com/supperdoggy/score/sctructs/service/items"
	"log"
)

func main() {
	// initializing db
	db, err := db3.InitItemsDB()
	if err != nil {
		panic(err.Error())
	}
	log.Println("successfully connected to db")
	defer db.Close()

	handlers := handlers2.Handlers{DB: db}

	r := gin.Default()

	apiv1 := r.Group(sctructs.ApiV1)
	{
		apiv1.POST(itemsdata.FindPath, handlers.Find)
		apiv1.POST(itemsdata.DeletePath, handlers.Delete)
		apiv1.POST(itemsdata.CreatePath, handlers.Create)
	}

	if err := r.Run(itemsdata.Port); err != nil {
		log.Println("!!! r.Run() ERROR !!!")
	}

}
