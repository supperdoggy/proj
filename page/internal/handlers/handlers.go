package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type obj map[string]interface{}

type Handlers struct {
}

func (h *Handlers) LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", obj{})
	return
}
