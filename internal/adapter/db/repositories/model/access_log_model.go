package model

import "time"

type AccessLog struct {
	ID         int `gorm:"primaryKey"`
	Path       string
	Method     string
	RemoteAddr string

	ReqQuery   string
	ReqBody    string
	ReqHeaders string

	CreatedAt time.Time `gorm:"index:idx_created_at"`
}
