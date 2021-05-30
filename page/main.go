package main

import (
	"github.com/gin-gonic/gin"
	handlers2 "github.com/supperdoggy/score/page/internal/handlers"
	"github.com/supperdoggy/score/sctructs"
	pagedata "github.com/supperdoggy/score/sctructs/service/page"
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

	auth := r.Group(pagedata.AuthPath)
	auth.Static("src/static", "./src/static")
	{
		auth.GET(pagedata.LoginPagePath, handlers.LoginPage)
		auth.GET(pagedata.RegisterPagePath, handlers.RegisterPage)
	}

	apiv1 := r.Group(sctructs.ApiV1)
	{
		apiv1.POST(pagedata.LoginReqPath, handlers.Login)
		apiv1.POST(pagedata.RegisterReqPath, handlers.Register)
	}

	main := r.Group("/", handlers.CheckToken)
	{
		main.GET("/", handlers.Index)
	}

	if err := r.Run(pagedata.Port); err != nil {
		log.Println("got error running application:", err.Error())
	}
}
