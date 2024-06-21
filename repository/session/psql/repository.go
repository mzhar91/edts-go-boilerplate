package psql

import (
	"context"

	uuid "github.com/satori/go.uuid"
	_model "sg-edts.com/edts-go-boilerplate/model"
	_repository "sg-edts.com/edts-go-boilerplate/pkg/repository"
)

type PsqlRepository interface {
	Read(ctx context.Context, conn *_repository.Use, param *_model.QuerySession) ([]*_model.Session, error)
	Total(ctx context.Context, conn *_repository.Use, param *_model.QuerySession) (int, error)
	ReadBy(ctx context.Context, conn *_repository.Use, param []*_model.ParamSession) ([]*_model.Session, error)
	Create(ctx context.Context, conn *_repository.Use, param any) error
	Update(ctx context.Context, conn *_repository.Use, id uuid.UUID, param *_model.Session) error
	Delete(ctx context.Context, conn *_repository.Use, id uuid.UUID) error
}
