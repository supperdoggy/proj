package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/supperdoggy/score/sctructs"
	"github.com/supperdoggy/score/sctructs/db"
	usersdata "github.com/supperdoggy/score/sctructs/service/users"
	"log"
	"net/http"
)

type obj map[string]interface{}
type arr []interface{}
type Handlers struct {
	DB db.IDB
}

// CreateUser - route to create new user, POST
func (h *Handlers) CreateUser(c *gin.Context) {
	var req usersdata.CreateUserRequest
	var res usersdata.CreateUserResponse
	if err := c.Bind(&req); err != nil {
		res.Error = errors.New("failed to parse your request")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if req.User.Username == "" || req.User.Age == 0 || req.User.Email == "" || req.User.HashedPass == "" {
		res.Error = errors.New("fill all fields")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := h.DB.Create(&req.User); err != nil {
		res.Error = errors.New("error creating new user")
		log.Println("CreateUser() -> Create(&user) error: ", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res.User = req.User
	c.JSON(http.StatusOK, res)
}

// GetAllUsers only for test, GET
func (h *Handlers) GetAllUsers(c *gin.Context) {
	var res usersdata.GetAllUsersResponse
	if dbresult := h.DB.Find(&res.Users); dbresult.Error != nil {
		res.Error = errors.New(dbresult.Error.Error())
		log.Println("GetAllUsers() -> Find() error:", dbresult.Error)
		c.JSON(http.StatusBadRequest, res)
	}

	c.JSON(http.StatusOK, res)

}

func (h *Handlers) Delete(c *gin.Context) {
	var req usersdata.DeleteRequest
	var res usersdata.DeleteResponse

	if err := c.Bind(&req); err != nil {
		res.Error = errors.New("failed to parse your request")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	// find by name
	if req.Username != "" {
		res.Error = h.DB.Where("Username = ?", req.Username).Delete(&res.User).Error
	} else if req.Email != "" {
		res.Error = h.DB.Where("Email = ?", req.Email).Delete(&res.User).Error
	} else {
		res.Error = h.DB.Find(req.ID).Delete(&res.User).Error
	}
	// else find by it

	if res.Error != nil {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handlers) Find(c *gin.Context) {
	var req usersdata.FindRequest
	var res usersdata.FindResponse
	if err := c.Bind(&req); err != nil {
		res.Error = errors.New("failed to parse your request")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	// find by name
	if req.Username != "" {
		res.Error = h.DB.Where("Username = ?", req.Username).First(&res.Users).Error
	} else if req.Email != "" {
		res.Error = h.DB.Where("Email = ?", req.Email).First(&res.Users).Error
	} else {
		res.Error = h.DB.First(&res.Users, req.ID).Error
	}
	// else find by it

	if res.Error != nil {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h Handlers) FindWithPassword(c *gin.Context) {
	var req usersdata.FindWithPasswordRequest
	var res usersdata.FindWithPasswordResponse

	if err := c.Bind(&req); err != nil {
		res.Error = errors.New("failed to parse your request")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	// find by name
	if req.Username != "" {
		res.Error = h.DB.Where("Username = ?", req.Username).First(&res.User).Error
	} else {
		res.Error = h.DB.Where("Email = ?", req.Email).First(&res.User).Error
	}
	// else find by it
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, res)
	}

	if req.Password != res.User.HashedPass {
		res.User = sctructs.User{}
		res.Error = errors.New("not found")
		c.JSON(http.StatusBadRequest, res)
	}
	c.JSON(http.StatusBadRequest, res)
}
