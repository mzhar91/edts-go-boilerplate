package psql

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/huandu/go-sqlbuilder"
	"github.com/sirupsen/logrus"

	uuid "github.com/satori/go.uuid"

	_constant "sg-edts.com/edts-go-boilerplate/constant"
	_helper "sg-edts.com/edts-go-boilerplate/helper"
	_model "sg-edts.com/edts-go-boilerplate/model"
	_repository "sg-edts.com/edts-go-boilerplate/pkg/repository"
)

type psqlRepository struct{}

func NewPsqlRepository() PsqlRepository {
	return &psqlRepository{}
}

func (m *psqlRepository) Read(ctx context.Context, conn *_repository.Use, param *_model.QuerySession) ([]*_model.Session, error) {
	if param == nil {
		param = &_model.QuerySession{}
	}

	offset := (param.Page - 1) * param.Limit
	sb := sqlbuilder.MySQL.NewSelectBuilder()
	sb.Select(
		"s.id AS id",
		"s.username AS username",
		"s.access_token AS access_token",
		"s.refresh_token AS refresh_token",
		"s.scope AS scope",
		"s.device_id AS device_id",
		"s.ip AS ip",
		"s.created_by AS created_by",
		"s.created_at AS created_at",
		"s.last_modified_by AS last_modified_by",
		"s.last_modified_at AS last_modified_at",
	)
	sb.From("session AS s")

	if param.Keyword != "" && param.FilterBy != "" {
		param.Keyword = strings.ToLower(param.Keyword)

		if param.FilterBy == _constant.FilterByUsername {
			sb.Where(
				sb.Equal("lower(s.username)", param.Keyword),
			)
		}
	}

	if param.SortBy != "" && param.OrderBy != "" {
		sb.OrderBy(_helper.CamelCaseToUnderscore(param.SortBy))

		switch param.OrderBy {
		case "asc":
			sb.Asc()
		case "desc":
			sb.Desc()
		default:
			sb.Desc()
		}
	} else {
		sb.OrderBy("a.created_at")
		sb.Desc()
	}

	if param.Page > 0 && param.Limit > 0 {
		sb.Limit(int(param.Limit)).Offset(int(offset))
	}

	query, args := sb.Build()

	result, err := m.fetchList(ctx, conn, query, args...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (m *psqlRepository) Total(ctx context.Context, conn *_repository.Use, param *_model.QuerySession) (int, error) {
	sb := sqlbuilder.MySQL.NewSelectBuilder()
	sb.Select("COUNT(1)")
	sb.From("session AS s")

	if param.Keyword != "" && param.FilterBy != "" {
		param.Keyword = strings.ToLower(param.Keyword)

		if param.FilterBy == _constant.FilterByUsername {
			sb.Where(
				sb.Equal("lower(s.username)", param.Keyword),
			)
		}
	}

	query, args := sb.Build()

	var row *sql.Row
	var count int

	if conn.Db != nil {
		row = conn.Db.QueryRowContext(ctx, query, args...)
	} else if conn.Trans != nil {
		row = conn.Trans.QueryRowContext(ctx, query, args...)
	}

	if row != nil {
		err := row.Scan(&count)
		if err != nil {
			return 0, err
		}

		return count, nil
	}

	return 0, nil
}

func (m *psqlRepository) ReadBy(ctx context.Context, conn *_repository.Use, param []*_model.ParamSession) ([]*_model.Session, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	sb.Select(
		"s.id AS id",
		"s.username AS username",
		"s.access_token AS access_token",
		"s.refresh_token AS refresh_token",
		"s.scope AS scope",
		"s.device_id AS device_id",
		"s.ip AS ip",
		"s.created_by AS created_by",
		"s.created_at AS created_at",
		"s.last_modified_by AS last_modified_by",
		"s.last_modified_at AS last_modified_at",
	)
	sb.From("session AS s")

	for _, v := range param {
		if v.Key != "" && v.Value != "" {
			var whereField string

			v.Value = strings.ToLower(v.Value)

			if v.Key == _constant.KeyID {
				whereField = "s.id"
			}

			if v.Key == _constant.KeyUsername {
				whereField = "s.username"
			}

			sb.Where(
				sb.Equal(whereField, v.Value),
			)
		}
	}

	query, args := sb.Build()

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
	ib.InsertInto("session")
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
		err = fmt.Errorf("Weird behaviour. Total affected: %d", affect)
		return err
	}

	return nil
}

func (m *psqlRepository) Update(ctx context.Context, conn *_repository.Use, id uuid.UUID, param *_model.Session) error {
	set, err := _repository.Set(reflect.ValueOf(param).Elem())
	if err != nil {
		return err
	}

	ub := sqlbuilder.PostgreSQL.NewUpdateBuilder()
	ub.Update("session")
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
		err = fmt.Errorf("Weird behaviour. Total affected: %d", affect)
		return err
	}

	return nil
}

func (m *psqlRepository) Delete(ctx context.Context, conn *_repository.Use, id uuid.UUID) error {
	db := sqlbuilder.PostgreSQL.NewDeleteBuilder()
	db.DeleteFrom("session")
	db.Where(
		db.Equal("id", id),
	)

	query, args := db.Build()

	stmt, err := conn.Trans.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affect == 0 {
		err = fmt.Errorf("Weird behaviour. Total affected: %d", affect)
		return err
	}

	return nil
}
