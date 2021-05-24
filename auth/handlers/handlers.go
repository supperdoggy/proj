package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/supperdoggy/score/auth/conf"
	"github.com/supperdoggy/score/auth/utils"
	"github.com/supperdoggy/score/sctructs"
	"github.com/supperdoggy/score/sctructs/communication"
	authdata "github.com/supperdoggy/score/sctructs/service/auth"
	usersdata "github.com/supperdoggy/score/sctructs/service/users"
	"log"
	"net/http"
	"regexp"
	"time"
)

type Handlers struct {
	Cache sctructs.AuthTokenCache
}

func (h *Handlers) CheckToken(c *gin.Context) {
	var req authdata.CheckTokenReq
	var res authdata.CheckTokenRes
	if err := c.Bind(&req); err != nil {
		res.Error = fmt.Sprintf("error binding your request")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	token, ok := h.Cache.Get(req.Token)
	if !ok || token.Expired() {
		c.JSON(http.StatusBadRequest, res)

		// if token expired we delete it from the cache
		if token.Expired() {
			h.Cache.Delete(token.Value)
		}
		return
	}
	res.Ok = true
	c.JSON(http.StatusOK, res)
}

// TODO: simplify
// жесть я нагамнякав, потом поправлю, честно честно
func (h *Handlers) Register(c *gin.Context) {
	var req authdata.RegisterReq
	var res authdata.RegisterRes
	if err := c.Bind(&req); err != nil {
		res.Error = fmt.Sprintf("error binding your request")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	// checking requirements for creds
	if err := utils.CheckCredentials(req); err != nil {
		res.Error = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}

	reqToUsers := usersdata.CreateUserRequest{User: sctructs.User{
		Username:   req.Username,
		Email:      req.Email,
		HashedPass: req.Password,
		BirthDate:  req.BirthDate,
	}}
	data, err := json.Marshal(reqToUsers)
	if err != nil {
		res.Error = fmt.Sprintf("error creating request to users " + err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	// creating user
	answer, err := communication.MakeHttpRequest(usersdata.UsersRoute+sctructs.ApiV1+usersdata.CreatePath, "POST", data)
	if err != nil {
		res.Error = fmt.Sprintf("error making request")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	var response usersdata.CreateUserResponse
	if err := json.Unmarshal(answer, &response); err != nil {
		res.Error = fmt.Sprintf("error unmarshal response from users " + string(answer))
		c.JSON(http.StatusBadRequest, res)
		return
	}

	// fix errors
	if response.Error != "" {
		res.Error = response.Error
		c.JSON(http.StatusBadRequest, res)
		return
	}

	// generating token
	token := sctructs.AuthToken{
		UserID:      response.User.ID,
		Username:    response.User.Username,
		Email:       response.User.Email,
		Value:       utils.GenerateRandomString(conf.TokenLen),
		TimeCreated: time.Now(),
		TimeExpired: time.Now().Add(64 * time.Hour),
	}
	h.Cache.Insert(token.Value, token)
	res.Token = token
	c.JSON(http.StatusOK, res)
}

func (h *Handlers) Login(c *gin.Context) {
	var req authdata.LoginReq
	var res authdata.LoginRes

	var reqToUsers usersdata.FindWithPasswordRequest
	var respFromUsers usersdata.FindWithPasswordResponse

	if err := c.Bind(&req); err != nil {
		res.Error = "Error binding your request"
		c.JSON(http.StatusBadRequest, res)
		return
	}
	// check if login is email
	if match, _ := regexp.MatchString(sctructs.RegexpEmail, req.Login); match {
		reqToUsers.Email = req.Login
	} else {
		reqToUsers.Username = req.Login
	}
	reqToUsers.Password = req.Password

	// sending req to users FindWithPasswordPath
	marshaledReq, err := json.Marshal(reqToUsers)
	if err != nil {
		res.Error = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}

	data, err := communication.MakeHttpRequest(usersdata.UsersRoute+usersdata.FindWithPasswordPath, "POST", marshaledReq)
	if err != nil {
		res.Error = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if err := json.Unmarshal(data, &respFromUsers); err != nil {
		res.Error = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}

	// check if found user
	if !respFromUsers.OK {
		res.Error = "Wrong password or login"
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if respFromUsers.Error != "" {
		log.Println("Login() -> respFromUsers.Error =", respFromUsers.Error)
	}

	// generating token
	token := sctructs.AuthToken{
		UserID:      respFromUsers.User.ID,
		Username:    respFromUsers.User.Username,
		Email:       respFromUsers.User.Email,
		Value:       utils.GenerateRandomString(conf.TokenLen),
		TimeCreated: time.Now(),
		TimeExpired: time.Now().Add(64 * time.Hour),
	}
	h.Cache.Insert(token.Value, token)
	res.Token = token
	res.OK = true
	c.JSON(http.StatusOK, res)
}
