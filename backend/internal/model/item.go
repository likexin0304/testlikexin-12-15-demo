package model

import (
	"time"

	"gorm.io/gorm"
)

// Item 示例数据模型
type Item struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	Description string         `json:"description"`
	Price       float64        `json:"price"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

