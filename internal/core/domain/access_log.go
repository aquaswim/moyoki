package domain

import "time"

type AccessLog struct {
	ID         int    `json:"id"`
	Path       string `json:"path"`
	Method     string `json:"method"`
	RemoteAddr string `json:"remoteAddr"`

	ReqQuery   string `json:"reqQuery"`
	ReqBody    string `json:"reqBody"`
	ReqHeaders string `json:"reqHeaders"`

	CreatedAt time.Time `json:"createdAt"`
}

type FindAccessLogParam struct {
	StartTime time.Time
	EndTime   time.Time
}
