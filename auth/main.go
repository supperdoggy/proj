package main

import (
	"github.com/gin-gonic/gin"
	authdata "github.com/supperdoggy/score/sctructs/service/auth"
	"log"
)

func main() {
	r := gin.Default()



	if err := r.Run(authdata.Port); err != nil {
		log.Println("error running service")
	}

}
