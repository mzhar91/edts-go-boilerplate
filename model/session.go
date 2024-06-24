package model

import (
	uuid "github.com/satori/go.uuid"
	"gopkg.in/guregu/null.v3"
)

type Session struct {
	ID           uuid.UUID   `json:"id" db:"id"`
	Username     string      `json:"username" db:"username"`
	Scope        string      `json:"scope" db:"scope"`
	DeviceID     string      `json:"device_id" db:"device_id"`
	CreatedDate  string      `json:"created_date" db:"created_date"`
	ModifiedDate null.String `json:"modified_date" db:"modified_date"`
}

type SessionResponse struct {
	ID           uuid.UUID   `json:"id"`
	Username     string      `json:"username"`
	Scope        string      `json:"scope"`
	DeviceID     string      `json:"device_id"`
	CreatedDate  string      `json:"created_date"`
	ModifiedDate null.String `json:"modified_date"`
}

type SessionSetLogin struct {
	ID           uuid.UUID   `json:"id" db:"id"`
	Username     string      `json:"username" db:"username"`
	Scope        string      `json:"scope" db:"scope"`
	DeviceID     string      `json:"device_id" db:"device_id"`
	CreatedDate  string      `json:"created_date" db:"created_date"`
	ModifiedDate null.String `json:"modified_date" db:"modified_date"`
}

type QuerySession struct {
	Page     int64  `json:"page" form:"page" query:"page"`
	Limit    int64  `json:"limit" form:"limit" query:"limit"`
	OrderBy  string `json:"order" form:"order" query:"order"`
	SortBy   string `json:"sort" form:"sort" query:"sort"`
	FilterBy string `json:"filter" form:"filter" query:"filter"`
	Keyword  string `json:"keyword" form:"keyword" query:"keyword"`
}

type ParamSession struct {
	Key   string `json:"key" form:"key" query:"key"`
	Value string `json:"value" form:"value" query:"value"`
}

type DropSession struct {
	SessionID   string `json:"session_id"`
	RequestInfo RequestInfo
}
