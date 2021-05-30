package main

import (
	"github.com/gin-gonic/gin"
	handlers2 "github.com/supperdoggy/score/auth/handlers"
	"github.com/supperdoggy/score/sctructs"
	authdata "github.com/supperdoggy/score/sctructs/service/auth"
	"log"
)

func main() {
	r := gin.Default()
	handlers := handlers2.Handlers{
		Cache:         sctructs.AuthTokenCache{}.Init(),
		UsernameCache: sctructs.AuthTokenCache{}.Init(),
	}

	apiv1 := r.Group(sctructs.ApiV1)
	{
		apiv1.POST(authdata.CheckTokenPath, handlers.CheckToken)
		apiv1.POST(authdata.RegisterPath, handlers.Register)
		apiv1.POST(authdata.LoginPath, handlers.Login)
		apiv1.POST(authdata.GetTokenByValuePath, handlers.GetTokenByValue)
	}

	if err := r.Run(authdata.Port); err != nil {
		log.Println("error running service")
	}

}
