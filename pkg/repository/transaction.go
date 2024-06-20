package repository

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	
	"github.com/sirupsen/logrus"
)

type Use struct {
	Db    *sql.DB
	Trans Transaction
}

type Transaction interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

type Trans func(Transaction) (error, int)

func WithTransaction(db *sql.DB, trans Trans) (err error, code int) {
	errCode := 0
	tx, err := db.Begin()
	if err != nil {
		return err, http.StatusInternalServerError
	}
	
	defer func() {
		if p := recover(); p != nil {
			e := tx.Rollback()
			if e != nil {
				logrus.Error(e)
				err = fmt.Errorf("rollback recover failed, please check your error log")
			}
			
			return
		} else if err != nil {
			e := tx.Rollback()
			if e != nil {
				logrus.Error(e)
				err = fmt.Errorf("rollback failed, please check your error log")
			}
			
			return
		} else {
			e := tx.Commit()
			if e != nil {
				logrus.Error(e)
				err = fmt.Errorf("commit failed, please check your error log")
			}
			
			return
		}
	}()
	
	err, errCode = trans(tx)
	
	return err, errCode
}
