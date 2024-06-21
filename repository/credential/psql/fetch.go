package psql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"

	_model "sg-edts.com/edts-go-boilerplate/model"
	_repository "sg-edts.com/edts-go-boilerplate/pkg/repository"
)

func (m *psqlRepository) fetchList(ctx context.Context, conn *_repository.Use, query string, args ...interface{}) (*_model.Credential, error) {
	var rows *sql.Rows
	var err error

	if conn.Db != nil {
		rows, err = conn.Db.QueryContext(ctx, query, args...)
	} else if conn.Trans != nil {
		rows, err = conn.Trans.QueryContext(ctx, query, args...)
	}

	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			logrus.Errorf("%v", err)
		}
	}(rows)

	for rows.Next() {
		t := new(_model.Credential)

		err = rows.Scan(
			&t.ID,
			&t.Username,
			&t.LastLogin,
			&t.CreatedDate,
			&t.ModifiedDate,
		)
		if err != nil {
			return nil, err
		}

		return t, nil
	}

	return nil, fmt.Errorf("data not found")
}

func (m *psqlRepository) fetchVerifyPassword(ctx context.Context, conn *_repository.Use, query string, args ...interface{}) (*_model.Credential, error) {
	var rows *sql.Rows
	var err error

	if conn.Db != nil {
		rows, err = conn.Db.QueryContext(ctx, query, args...)
	} else if conn.Trans != nil {
		rows, err = conn.Trans.QueryContext(ctx, query, args...)
	}

	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			logrus.Errorf("%v", err)
		}
	}(rows)

	for rows.Next() {
		t := new(_model.Credential)

		err = rows.Scan(
			&t.ID,
			&t.Username,
		)
		if err != nil {
			return nil, err
		}

		return t, nil
	}

	return nil, fmt.Errorf("data not found")
}
