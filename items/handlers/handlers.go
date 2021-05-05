package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/supperdoggy/score/sctructs"
	"github.com/supperdoggy/score/sctructs/db"
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

func (h *Handlers) Create(c *gin.Context) {
	var req sctructs.Item

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, obj{"error": "error binding request"})
		return
	}

	if req.Author == "" || req.Category == "" || req.Name == "" {
		c.JSON(http.StatusBadRequest, obj{"error": "need to fill required fields"})
		return
	}

	if err := h.DB.Create(&req); err != nil {
		c.JSON(http.StatusBadRequest, obj{"error": "error creating item in db"})
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handlers) Delete(c *gin.Context) {
	var req struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, obj{"error": "binding error"})
		return
	}

	var result sctructs.Item
	var err error
	// find by name
	if req.Name != "" {
		err = h.DB.Where("Name = ?", req.Name).Delete(&result).Error
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
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Author string `json:"author"`
	}

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, obj{"error": "binding error"})
		return
	}

	var result []sctructs.Item
	var err error
	// find by name
	if req.Name != "" {
		err = h.DB.Where("Name = ?", req.Name).Find(&result).Error
	} else if req.Author != "" {
		err = h.DB.Where("Author = ?", req.Author).Find(&result).Error
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
