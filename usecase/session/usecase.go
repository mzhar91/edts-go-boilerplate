package session

import (
	"context"
	
	_model "sg-edts.com/edts-go-boilerplate/model"
	_auth "sg-edts.com/edts-go-boilerplate/pkg/auth"
)

type Usecase interface {
	GetOwnSession(ctx context.Context, claims *_auth.Claims) ([]*_model.SessionResponse, error)
	GetSession(ctx context.Context, param *_model.QuerySession) ([]*_model.SessionResponse, int, error)
	DropOwnSession(ctx context.Context, claims *_auth.Claims, param *_model.DropSession) (res string, err error)
	DropSession(ctx context.Context, claims *_auth.Claims, param *_model.DropSession) (res string, err error)
}
