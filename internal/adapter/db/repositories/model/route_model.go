package model

import (
	"gorm.io/datatypes"
	"time"
)

type Route struct {
	ID          int    `gorm:"primaryKey"`
	Path        string `gorm:"uniqueIndex"`
	Description string

	ResponseCode    int
	ResponseBody    string
	ResponseHeaders datatypes.JSON

	CreatedAt time.Time
	UpdatedAt time.Time
}
