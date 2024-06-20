package credential

import (
	"context"
	
	uuid "github.com/satori/go.uuid"
	_model "sg-edts.com/edts-go-boilerplate/model"
	_repository "sg-edts.com/edts-go-boilerplate/pkg/repository"
)

type PsqlRepository interface {
	GetByUsername(ctx context.Context, conn *_repository.Use, username string) (*_model.Credential, error)
	VerifyPassword(ctx context.Context, conn *_repository.Use, username string, password string) (*_model.Credential, error)
	Create(ctx context.Context, conn *_repository.Use, param any) error
	Update(ctx context.Context, conn *_repository.Use, id uuid.UUID, param any) error
	CheckUsername(ctx context.Context, conn *_repository.Use, username string) (bool, error)
}
