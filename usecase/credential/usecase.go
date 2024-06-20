package credential

import (
	"context"
	
	_model "sg-edts.com/edts-go-boilerplate/model"
	_auth "sg-edts.com/edts-go-boilerplate/pkg/auth"
)

type Usecase interface {
	RefreshToken(ctx context.Context, app string, cCtx *_auth.ClaimsContext) (*_model.SignInResponse, error, map[string]interface{})
	SignIn(ctx context.Context, app string, cCtx *_auth.ClaimsContext, req *_model.SignInRequest) (*_model.SignInResponse, error, map[string]interface{})
	SignOut(ctx context.Context, app string, email string, req *_model.SignOutRequest) (error, map[string]interface{})
	AddCredential(ctx context.Context, app string, claims *_auth.Claims, req *_model.AddCredentialRequest) (*_model.AddCredentialResponse, error, map[string]interface{})
}
