package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/supperdoggy/score/sctructs/db"
	itemsdata "github.com/supperdoggy/score/sctructs/service/items"
	"net/http"
)

type obj map[string]interface{}
type arr []interface{}
type Handlers struct {
	DB db.IDB
}

func (h *Handlers) HelloWorld(c *gin.Context) {
	c.JSON(200, obj{"hello": "world"})
}

// Create - handler for creating new item in db
func (h *Handlers) Create(c *gin.Context) {
	var req itemsdata.CreateRequest
	var res itemsdata.CreateResponse
	if err := c.Bind(&req); err != nil {
		res.Error = errors.New("error binging your request")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if req.Item.Author == "" || req.Item.Category == "" || req.Item.Name == "" {
		res.Error = errors.New("need to fill required fields")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := h.DB.Create(&req.Item); err != nil {
		res.Error = errors.New("error creating item in db")
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
		res.Error = errors.New("binding error")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	// find by name
	if req.Name != "" {
		res.Error = h.DB.Where("Name = ? AND Author = ?", req.Name, req.Author).Delete(&res.Item).Error
	} else {
		res.Error = h.DB.Find(req.ID).Delete(&res.Item).Error
	}
	// else find by it

	if res.Error != nil {
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
		res.Error = errors.New("binding error")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	// find by name
	if req.Name != "" {
		res.Error = h.DB.Where("Name = ? AND Author = ?", req.Name, req.Author).First(&res.Item).Error
	} else {
		res.Error = h.DB.First(&res.Item, req.ID).Error
	}
	// else find by it

	if res.Error != nil {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
