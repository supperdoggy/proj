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
	c.Status(http.StatusOK)
}

// GetAllUsers only for test, GET
func (h *Handlers) GetAllUsers(c *gin.Context) {
	var users []sctructs.User
	// how to search
	thing := h.DB.Where("Username = ?", "username random")
	if dbresult := h.DB.Find(&users); dbresult.Error != nil {
		log.Println("GetAllUsers() -> Find() error:", dbresult.Error)
		c.JSON(http.StatusBadRequest, obj{"error": dbresult.Error})
	}
	thing.First(&users)

	c.JSON(http.StatusOK, users)

}
