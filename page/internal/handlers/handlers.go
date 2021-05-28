package handlers

import (
	"encoding/json"
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

	req.Login = c.PostForm("login")
	req.Password = c.PostForm("password")
	if req.Login == "" || req.Password  == "" {
		resp.Error = "you need to fill login and password fields"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	data, err := authapi.ApiV1(authdata.LoginPath, "POST", req)
	if err != nil {
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	var resultFromAuth authdata.LoginRes
	if err := json.Unmarshal(data, &resultFromAuth); err != nil {
		resp.Error = "error unmarshaling"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if resultFromAuth.Error != "" {
		log.Println("got error from auth login", resultFromAuth.Error)
	}
	log.Println(resultFromAuth)
	if !resultFromAuth.OK {
		resp.Error = "wrong login or password"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.Token = resultFromAuth.Token
	resp.OK = true

	c.JSON(http.StatusOK, resp)
}

func (h *Handlers) LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", obj{})
	return
}
