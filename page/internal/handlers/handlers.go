package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/supperdoggy/score/page/internal/authapi"
	authdata "github.com/supperdoggy/score/sctructs/service/auth"
	pagedata "github.com/supperdoggy/score/sctructs/service/page"
	"log"
	"net/http"
)

type obj map[string]interface{}

type Handlers struct {
}

func (h *Handlers) Login(c *gin.Context) {
	var req pagedata.LoginRequest
	var resp pagedata.LoginResponse

	if err := c.Bind(&req); err != nil {
		resp.Error = "binding error " + err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if req.Login == "" || req.Password  == "" {
		resp.Error = "you need to fill login and password fields"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resultFromAuth, err := authapi.ApiV1(authdata.LoginPath, "POST", req)
	if err != nil {
		resp.Error = err.Error()
		return
	}
	if resultFromAuth.(authdata.LoginRes).Error != "" {
		log.Println("got error from auth login", resultFromAuth.(authdata.LoginRes).Error)
	}
	if !resultFromAuth.(authdata.LoginRes).OK {
		resp.Error = "wrong login or password"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.Token = resultFromAuth.(authdata.LoginRes).Token
	resp.OK = true

	c.JSON(http.StatusOK, resp)
}

func (h *Handlers) LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", obj{})
	return
}
