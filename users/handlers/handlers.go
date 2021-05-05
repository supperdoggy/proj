package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/supperdoggy/score/sctructs"
	"github.com/supperdoggy/score/sctructs/db"
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
	var user sctructs.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, obj{"error": "failed to parse your request"})
		return
	}

	if user.Username == "" || user.Age == 0 || user.Email == "" || user.HashedPass == "" {
		c.JSON(http.StatusBadRequest, obj{"error": "fill all fields"})
		return
	}

	if err := h.DB.Create(&user); err != nil {
		log.Println("CreateUser() -> Create(&user) error: ", err.Error())
		c.JSON(http.StatusBadRequest, obj{"error": "error creating new user"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// GetAllUsers only for test, GET
func (h *Handlers) GetAllUsers(c *gin.Context) {
	var users []sctructs.User
	// how to search
	if dbresult := h.DB.Find(&users); dbresult.Error != nil {
		log.Println("GetAllUsers() -> Find() error:", dbresult.Error)
		c.JSON(http.StatusBadRequest, obj{"error": dbresult.Error})
	}

	c.JSON(http.StatusOK, users)

}

func (h *Handlers) Delete(c *gin.Context) {
	var req struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, obj{"error": "binding error"})
		return
	}

	var result sctructs.User
	var err error
	// find by name
	if req.Username != "" {
		err = h.DB.Where("Username = ?", req.Username).Delete(&result).Error
	} else if req.Email != "" {
		err = h.DB.Where("Email = ?", req.Email).Delete(&result).Error
	} else {
		err = h.DB.Find(req.ID).Delete(&result).Error
	}
	// else find by it

	if err != nil {
		c.JSON(http.StatusBadRequest, obj{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *Handlers) Find(c *gin.Context) {
	var req struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, obj{"error": "binding error"})
		return
	}

	var result []sctructs.Item
	var err error
	// find by name
	if req.Username != "" {
		err = h.DB.Where("Username = ?", req.Username).Find(&result).Error
	} else if req.Email != "" {
		err = h.DB.Where("Email = ?", req.Email).Find(&result).Error
	} else {
		err = h.DB.First(&result, req.ID).Error
	}
	// else find by it

	if err != nil {
		c.JSON(http.StatusBadRequest, obj{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
