package psql

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"

	"github.com/huandu/go-sqlbuilder"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	_model "sg-edts.com/edts-go-boilerplate/model"
	_repository "sg-edts.com/edts-go-boilerplate/pkg/repository"
)

type psqlRepository struct{}

func NewPsqlRepository() PsqlRepository {
	return &psqlRepository{}
}

func (m *psqlRepository) VerifyPassword(ctx context.Context, conn *_repository.Use, username string, password string) (*_model.Credential, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	sb.Select(
		"u.id",
		"u.username",
	)
	sb.From("user AS u")
	sb.Where(
		sb.And(
			sb.Equal("c.username", username),
			sb.Equal("c.password", password),
		),
	)

	query, args := sb.Build()

	result, err := m.fetchVerifyPassword(ctx, conn, query, args...)
	if err != nil {
		return nil, fmt.Errorf("password didn't match")
	}

	return result, nil
}

func (m *psqlRepository) CheckUsername(ctx context.Context, conn *_repository.Use, username string) (bool, error) {
	var (
		count int
		row   *sql.Row
	)

	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	sb.Select("COUNT(1)")
	sb.From("user AS u")
	sb.Where(
		sb.Equal("u.username", username),
	)

	query, args := sb.Build()

	if conn.Db != nil {
		row = conn.Db.QueryRowContext(ctx, query, args...)
	} else if conn.Trans != nil {
		row = conn.Trans.QueryRowContext(ctx, query, args...)
	}

	if row != nil {
		err := row.Scan(&count)
		if err != nil {
			return false, err
		}

		if count > 0 {
			return true, err
		}
	}

	return false, nil
}

func (m *psqlRepository) GetByUsername(ctx context.Context, conn *_repository.Use, username string) (*_model.Credential, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	sb.Select(
		"c.id AS id",
		"c.username AS username",
		"c.last_login AS last_login",
		"c.created_date AS created_date",
		"c.modified_date AS modified_date",
	)
	sb.From("user AS u")
	sb.Where(
		sb.Equal("u.username", username),
	)

	query, args := sb.Build()

	logrus.Infof("GetByUsername %v : %v", query, args)

	result, err := m.fetchList(ctx, conn, query, args...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (m *psqlRepository) Create(ctx context.Context, conn *_repository.Use, param any) error {
	cols, values, err := _repository.Values(reflect.ValueOf(param).Elem())
	if err != nil {
		return err
	}

	ib := sqlbuilder.PostgreSQL.NewInsertBuilder()
	ib.InsertInto("user")
	ib.Cols(cols...)
	ib.Values(values...)

	query, args := ib.Build()

	stmt, err := conn.Trans.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			logrus.Errorf("%v", err)
		}
	}(stmt)

	res, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affect == 0 {
		err = fmt.Errorf("weird behaviour. Total affected: %d", affect)
		return err
	}

	return nil
}

func (m *psqlRepository) Update(ctx context.Context, conn *_repository.Use, id uuid.UUID, param any) error {
	set, err := _repository.Set(reflect.ValueOf(param).Elem())
	if err != nil {
		return err
	}

	ub := sqlbuilder.PostgreSQL.NewUpdateBuilder()
	ub.Update("user")
	ub.Set(set...)
	ub.Where(
		ub.Equal("id", id),
	)

	query, args := ub.Build()

	stmt, err := conn.Trans.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			logrus.Errorf("%v", err)
		}
	}(stmt)

	res, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affect == 0 {
		err = fmt.Errorf("weird behaviour. Total affected: %d", affect)
		return err
	}

	return nil
}
