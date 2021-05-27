package main

import (
	"github.com/gin-gonic/gin"
	handlers2 "github.com/supperdoggy/score/page/internal/handlers"
	"log"
)

func init() {
	handlers = handlers2.Handlers{}
}

var handlers handlers2.Handlers

func main() {
	r := gin.Default()
	r.Static("src/static", "./src/static")
	r.LoadHTMLGlob("src/template/*")

	auth := r.Group("/auth")
	auth.Static("src/static", "./src/static")
	{
		auth.GET("/login", handlers.LoginPage)
	}

	if err := r.Run(":12212"); err != nil {
		log.Println("got error running application:", err.Error())
	}
}
