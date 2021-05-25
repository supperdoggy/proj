package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/supperdoggy/score/sctructs/db"
	itemsdata "github.com/supperdoggy/score/sctructs/service/items"
	"gorm.io/gorm"
	"net/http"
)

type obj map[string]interface{}
type arr []interface{}
type Handlers struct {
	DB db.IDB
}

// Create - handler for creating new item in db
func (h *Handlers) Create(c *gin.Context) {
	var req itemsdata.CreateRequest
	var res itemsdata.CreateResponse
	if err := c.Bind(&req); err != nil {
		res.Error = fmt.Sprintf("error binging your request")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if req.Item.Author == "" || req.Item.Category == "" || req.Item.Name == "" {
		res.Error = fmt.Sprintf("need to fill required fields")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := h.DB.Create(&req.Item); err != nil {
		res.Error = fmt.Sprintf("error creating item in db")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res.Item = req.Item
	c.JSON(http.StatusOK, res)
}

// Delete - handler for deleting item from db
func (h *Handlers) Delete(c *gin.Context) {
	var req itemsdata.DeleteRequest
	var res itemsdata.DeleteResponse

	if err := c.Bind(&req); err != nil {
		res.Error = fmt.Sprintf("binding error")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	// find by name
	var result *gorm.DB
	if req.Name != "" {
		result = h.DB.Where("Name = ? AND Author = ?", req.Name, req.Author).Delete(&res.Item)
	} else {
		result = h.DB.Find(req.ID).Delete(&res.Item)
	}
	if result != nil && result.Error != nil {
		res.Error = result.Error.Error()
	}
	// else find by it

	if res.Error != "" {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

// Find - handler for finding specific item in db
func (h *Handlers) Find(c *gin.Context) {
	var req itemsdata.FindRequest
	var res itemsdata.FindResponse
	if err := c.Bind(&req); err != nil {
		res.Error = fmt.Sprintf("binding error")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	// find by name
	var result *gorm.DB
	if req.Name != "" {
		result = h.DB.Where("Name = ? AND Author = ?", req.Name, req.Author).First(&res.Item)
	} else {
		result = h.DB.First(&res.Item, req.ID)
	}
	if result != nil && result.Error != nil {
		res.Error = result.Error.Error()
	}
	// else find by it

	if res.Error != "" {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
