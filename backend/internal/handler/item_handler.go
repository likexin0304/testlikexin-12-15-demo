package handler

import (
	"net/http"
	"strconv"

	"{{.ProjectName}}/internal/model"
	"{{.ProjectName}}/pkg/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ItemHandler struct {
	db *gorm.DB
}

func NewItemHandler(db *gorm.DB) *ItemHandler {
	return &ItemHandler{db: db}
}

// GetItems 获取所有物品
func (h *ItemHandler) GetItems(c *gin.Context) {
	var items []model.Item
	if err := h.db.Find(&items).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to fetch items")
		return
	}

	response.Success(c, gin.H{
		"items": items,
		"total": len(items),
	})
}

// GetItem 获取单个物品
func (h *ItemHandler) GetItem(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	var item model.Item
	if err := h.db.First(&item, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			response.Error(c, http.StatusNotFound, "Item not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, "Failed to fetch item")
		return
	}

	response.Success(c, item)
}

// CreateItem 创建物品
func (h *ItemHandler) CreateItem(c *gin.Context) {
	var item model.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.db.Create(&item).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create item")
		return
	}

	response.Success(c, item)
}

// UpdateItem 更新物品
func (h *ItemHandler) UpdateItem(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	var item model.Item
	if err := h.db.First(&item, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			response.Error(c, http.StatusNotFound, "Item not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, "Failed to fetch item")
		return
	}

	if err := c.ShouldBindJSON(&item); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.db.Save(&item).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to update item")
		return
	}

	response.Success(c, item)
}

// DeleteItem 删除物品
func (h *ItemHandler) DeleteItem(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := h.db.Delete(&model.Item{}, id).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to delete item")
		return
	}

	response.Success(c, gin.H{"message": "Item deleted successfully"})
}

