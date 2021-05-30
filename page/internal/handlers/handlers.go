package handlers

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/supperdoggy/score/page/internal/authapi"
	"github.com/supperdoggy/score/page/internal/utils"
	authdata "github.com/supperdoggy/score/sctructs/service/auth"
	pagedata "github.com/supperdoggy/score/sctructs/service/page"
	"log"
	"net/http"
)

type obj map[string]interface{}

type Handlers struct {
}

func (h *Handlers) Register(c *gin.Context) {
	var req pagedata.RegisterRequest
	var resp pagedata.RegisterResponse

	req.Username = c.PostForm("username")
	req.Password = c.PostForm("password")
	req.Name = c.PostForm("name")
	req.Email = c.PostForm("email")

	if req.Username == "" || req.Password == "" || req.Email == "" || req.Name == "" {
		resp.Error = "fill all the fields"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	data, err := authapi.ApiV1(authdata.RegisterPath, "POST", req)
	if err != nil {
		resp.Error = "error making request to auth"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	var resultFromAuth authdata.RegisterRes
	if err := json.Unmarshal(data, &resultFromAuth); err != nil {
		resp.Error = "error getting response from auth"
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	if resultFromAuth.Error != "" {
		resp.Error = resultFromAuth.Error
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	resp.Token = resultFromAuth.Token
	resp.OK = true
	c.SetCookie("t", resp.Token.Value, 999999, "/", "localhost", false, false)
	c.JSON(http.StatusOK, resp)
}

func (h *Handlers) Login(c *gin.Context) {
	var req pagedata.LoginRequest
	var resp pagedata.LoginResponse

	req.Login = c.PostForm("login")
	req.Password = c.PostForm("password")
	if req.Login == "" || req.Password == "" {
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
	if !resultFromAuth.OK {
		resp.Error = "wrong login or password"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.Token = resultFromAuth.Token
	resp.OK = true

	c.SetCookie("t", resp.Token.Value, 999999, "/", "localhost", false, false)

	c.JSON(http.StatusOK, resp)
}

func (h *Handlers) CheckToken(c *gin.Context) {
	token, err := c.Cookie("t")
	if err != nil {
		defer c.Redirect(http.StatusMovedPermanently, "/auth/login")
		err := c.AbortWithError(http.StatusForbidden, errors.New("no token"))
		log.Println("checkToken", err.Err)
		return
	}
	var checkreq authdata.CheckTokenReq
	var checkresp authdata.CheckTokenRes
	checkreq.Token = token

	data, err := authapi.ApiV1(authdata.CheckTokenPath, "POST", checkreq)
	if err := json.Unmarshal(data, &checkresp); err != nil || !checkresp.Ok {
		defer c.Redirect(http.StatusMovedPermanently, "/auth/login")
		err = c.AbortWithError(http.StatusForbidden, errors.New("cookie expired"))
		log.Println("checkToken", err.Error())
		c.SetCookie("t", "", -1, "/", "localhost", false, false)
		return
	}
}

func (h *Handlers) LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", obj{})
}

func (h *Handlers) RegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", obj{})
}

func (h *Handlers) Index(c *gin.Context) {
	// todo get user struct
	var username string
	token, err := utils.GetAuthToken(c)
	if err != nil {
		log.Println("error getting token struct")
	} else {
		username = token.Username
	}

	//var req itemsdata.

	c.HTML(http.StatusOK, "index.html", obj{"username": username})
}
