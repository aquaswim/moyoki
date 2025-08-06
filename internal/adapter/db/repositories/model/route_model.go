package model

import (
	"gorm.io/datatypes"
	"time"
)

type Route struct {
	ID          int    `gorm:"primaryKey"`
	Method      string `gorm:"index:,unique,composite:idx_method_path"`
	Path        string `gorm:"index:,unique,composite:idx_method_path"`
	Description string

	ResponseCode    int
	ResponseBody    string
	ResponseHeaders datatypes.JSON

	CreatedAt time.Time
	UpdatedAt time.Time
}
