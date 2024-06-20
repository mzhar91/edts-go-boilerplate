package model

import (
	uuid "github.com/satori/go.uuid"
	"gopkg.in/guregu/null.v3"
)

type Credential struct {
	ID           uuid.UUID   `db:"id"`
	Username     string      `db:"username"`
	UserID       uuid.UUID   `db:"user_id"`
	PasswordHash string      `db:"password_hash"`
	Scope        string      `db:"scope"`
	LastLogin    string      `db:"last_login"`
	CreatedDate  string      `db:"created_date"`
	ModifiedDate null.String `db:"modified_date"`
}

type CredentialSetLogIn struct {
	LastLogin    string `db:"last_login"`
	ModifiedDate string `db:"modified_date"`
}

type CredentialSetLogOut struct {
	ModifiedDate string `json:"modified_date" db:"modified_date"`
}

type CredentialResponse struct {
	ID           uuid.UUID   `json:"id"`
	Username     string      `json:"username"`
	LastLogin    string      `json:"last_login"`
	CreatedDate  string      `json:"created_date"`
	ModifiedDate null.String `json:"modified_date"`
}

type QueryCredential struct {
	Page    int64  `json:"page" form:"page" query:"page"`
	Limit   int64  `json:"limit" form:"limit" query:"limit"`
	OrderBy string `json:"order" form:"order" query:"order"`
	SortBy  string `json:"sort" form:"sort" query:"sort"`
	Keyword string `json:"keyword" form:"keyword" query:"keyword"`
}

type AddCredentialRequest struct {
	Username    string `json:"username" validate:"required,email"`
	Scope       string `json:"scope" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Claim       string `json:"claim" validate:"required"`
	RequestInfo RequestInfo
}

type AddCredentialResponse struct {
	Username string `json:"username"`
}

type SignInResponse struct {
	AccessToken     string `json:"access_token"`
	RefreshToken    string `json:"refresh_token"`
	TokenExpiration int64  `json:"token_exp"`
}

type TokenResponse struct {
	AccessToken     string `json:"access_token"`
	RefreshToken    string `json:"refresh_token"`
	TokenExpiration int64  `json:"token_exp"`
}

type SignInRequest struct {
	Username    string `json:"username" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
	DeviceId    string `json:"device_id" validate:"required"`
	RequestInfo RequestInfo
}

type ExternalSignInRequest struct {
	Email       string `json:"email" validate:"required,email"`
	ProjectID   string `json:"project_id" validate:"required"`
	RequestInfo RequestInfo
}

type SignOutRequest struct {
	DeviceId    string `json:"device_id" validate:"required"`
	RequestInfo RequestInfo
}

type VerifyResetPassRequest struct {
	Token              string `json:"token" validate:"required"`
	NewPassword        string `json:"password" validate:"required"`
	ConfirmNewPassword string `json:"confirm_password" validate:"required"`
	RequestInfo        RequestInfo
}

type VerifyChangePassRequest struct {
	Token              string `json:"token" validate:"required"`
	NewPassword        string `json:"password" validate:"required"`
	ConfirmNewPassword string `json:"confirm_password" validate:"required"`
	RequestInfo        RequestInfo
}

type ResetPasswordRequest struct {
	Username    string `json:"username" validate:"required,email"`
	RequestInfo RequestInfo
}

type ConfirmResetPasswordRequest struct {
	Token       string `json:"token" validate:"required"`
	Password    string `json:"password" validate:"required"`
	RequestInfo RequestInfo
}

type ValidateResetPasswordRequest struct {
	Token       string `json:"token" validate:"required"`
	RequestInfo RequestInfo
}

type ChangePasswordRequest struct {
	Email       string `json:"email" validate:"required,email"`
	RequestInfo RequestInfo
}
