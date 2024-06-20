package model

import (
	uuid "github.com/satori/go.uuid"
	"gopkg.in/guregu/null.v3"
)

type Session struct {
	ID             uuid.UUID   `json:"id" db:"id"`
	Username       string      `json:"username" db:"username"`
	AccessToken    string      `json:"access_token" db:"access_token"`
	RefreshToken   string      `json:"refresh_token" db:"refresh_token"`
	Scope          string      `json:"scope" db:"scope"`
	DeviceID       string      `json:"device_id" db:"device_id"`
	Ip             string      `json:"ip" db:"ip"`
	Host           string      `json:"host" db:"host"`
	UserAgent      string      `json:"user_agent" db:"user_agent"`
	CreatedAt      int64       `json:"created_at" db:"created_at"`
	CreatedBy      string      `json:"created_by" db:"created_by"`
	LastModifiedBy null.String `json:"last_modified_by" db:"last_modified_by"`
	LastModifiedAt null.Int    `json:"last_modified_at" db:"last_modified_at"`
}

type SessionResponse struct {
	ID             uuid.UUID   `json:"id"`
	Username       string      `json:"username"`
	AccessToken    string      `json:"access_token"`
	RefreshToken   string      `json:"refresh_token"`
	Scope          string      `json:"scope"`
	DeviceID       string      `json:"device_id"`
	Ip             string      `json:"ip"`
	Host           string      `json:"host"`
	UserAgent      string      `json:"user_agent"`
	CreatedAt      int64       `json:"created_at"`
	CreatedBy      string      `json:"created_by"`
	LastModifiedBy null.String `json:"last_modified_by"`
	LastModifiedAt null.Int    `json:"last_modified_at"`
}

type SessionSetLogin struct {
	Username     string `json:"username" db:"username"`
	AccessToken  string `json:"access_token" db:"access_token"`
	RefreshToken string `json:"refresh_token" db:"refresh_token"`
	Scope        string `json:"scope" db:"scope"`
	DeviceID     string `json:"device_id" db:"device_id"`
	Ip           string `json:"ip" db:"ip"`
	Host         string `json:"host" db:"host"`
	UserAgent    string `json:"user_agent" db:"user_agent"`
	CreatedAt    int64  `json:"created_at" db:"created_at"`
	CreatedBy    string `json:"created_by" db:"created_by"`
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
