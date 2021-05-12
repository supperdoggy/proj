package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/supperdoggy/score/sctructs"
	"github.com/supperdoggy/score/sctructs/communication"
	authdata "github.com/supperdoggy/score/sctructs/service/auth"
	usersdata "github.com/supperdoggy/score/sctructs/service/users"
	"net/http"
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
		return
	}
	res.Ok = true
	c.JSON(http.StatusOK, res)
}

// TODO: simplify
func (h *Handlers) Register(c *gin.Context) {
	var req authdata.RegisterReq
	var res authdata.RegisterRes
	if err := c.Bind(&req); err != nil {
		res.Error = fmt.Sprintf("error binding your request")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if req.Username == "" || req.Email == "" || req.Password == "" {
		res.Error = fmt.Sprintf("fill all of the fields")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	reqToUsers := usersdata.CreateUserRequest{User: sctructs.User{
		Username:     req.Username,
		Email:        req.Email,
		HashedPass:   req.Password,
		BirthDate:    req.BirthDate,
	}}
	data, err := json.Marshal(reqToUsers)
	if err != nil {
		res.Error = fmt.Sprintf("error creating request to users " + err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	// creating user
	answer, err := communication.MakeHttpRequest("http://"+usersdata.UsersRoute+sctructs.ApiV1+usersdata.CreatePath, "POST", data)
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
		// TODO: create random string generator
		Value:       "adsadsads",
		TimeCreated: time.Now(),
		TimeExpired: time.Now().Add(64*time.Hour),
	}
	h.Cache.Insert(token.Value, token)
	res.Token = token
	c.JSON(http.StatusOK, res)
}
