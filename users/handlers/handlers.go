package handlers

import (
	"fmt"
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
		res.Error = fmt.Sprintf("failed to parse your request")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if req.User.Username == "" || req.User.Email == "" || req.User.HashedPass == "" {
		res.Error = fmt.Sprintf("fill all fields")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	userExists := sctructs.User{}
	err1 := h.DB.Where("Username = ?", req.User.Username).First(&userExists).Error
	err1 = h.DB.Where("Email = ?", req.User.Email).First(&userExists).Error
	if err1 == nil {
		res.Error = fmt.Sprintf("user with given username or email already exists")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := h.DB.Create(&req.User); err != nil {
		res.Error = fmt.Sprintf("error creating new user")
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
		res.Error = fmt.Sprintf(dbresult.Error.Error())
		log.Println("GetAllUsers() -> Find() error:", dbresult.Error)
		c.JSON(http.StatusBadRequest, res)
	}

	c.JSON(http.StatusOK, res)

}

// Delete - handler for deleting user from db
func (h *Handlers) Delete(c *gin.Context) {
	var req usersdata.DeleteRequest
	var res usersdata.DeleteResponse

	if err := c.Bind(&req); err != nil {
		res.Error = fmt.Sprintf("failed to parse your request")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	// find by name
	if req.Username != "" {
		res.Error = h.DB.Where("Username = ?", req.Username).Delete(&res.User).Error.Error()
	} else if req.Email != "" {
		res.Error = h.DB.Where("Email = ?", req.Email).Delete(&res.User).Error.Error()
	} else {
		res.Error = h.DB.Find(req.ID).Delete(&res.User).Error.Error()
	}
	// else find by it

	if res.Error != "" {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

// Find - handler for finding specific user in db
func (h *Handlers) Find(c *gin.Context) {
	var req usersdata.FindRequest
	var res usersdata.FindResponse
	if err := c.Bind(&req); err != nil {
		res.Error = fmt.Sprintf("failed to parse your request")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	// find by name
	if req.Username != "" {
		res.Error = h.DB.Where("Username = ?", req.Username).First(&res.Users).Error.Error()
	} else if req.Email != "" {
		res.Error = h.DB.Where("Email = ?", req.Email).First(&res.Users).Error.Error()
	} else {
		res.Error = h.DB.First(&res.Users, req.ID).Error.Error()
	}
	// else find by it

	if res.Error != "" {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

// FindWithPassword - a tricky one, find a specific user and then
// if given password is equal to found user, returning that user
func (h Handlers) FindWithPassword(c *gin.Context) {
	var req usersdata.FindWithPasswordRequest
	var res usersdata.FindWithPasswordResponse

	if err := c.Bind(&req); err != nil {
		res.Error = fmt.Sprintf("failed to parse your request")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	// find by name
	if req.Username != "" {
		res.Error = h.DB.Where("Username = ?", req.Username).First(&res.User).Error.Error()
	} else {
		res.Error = h.DB.Where("Email = ?", req.Email).First(&res.User).Error.Error()
	}
	// else find by it
	if res.Error != "" {
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if req.Password != res.User.HashedPass {
		res.User = sctructs.User{}
		res.Error = fmt.Sprintf("not found")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusBadRequest, res)
}
