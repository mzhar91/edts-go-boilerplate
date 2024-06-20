package session

import (
	"fmt"
	
	"github.com/huandu/go-sqlbuilder"
)

func (m *psqlRepository) tokenSelectCondition(sql *sqlbuilder.SelectBuilder, token string, tipe string) bool {
	sql.Where(
		sql.Equal(fmt.Sprintf("c.%v", tipe), token),
	)
	
	// select condition
	switch tipe {
	case "reset_password_token":
		sql.Where(
			sql.Equal("c.is_reset_password_verified", true),
		)
	case "change_password_token":
		sql.Where(
			sql.Equal("c.is_locked", false),
			sql.Equal("c.is_change_password_verified", false),
		)
	default:
		return false
	}
	
	return true
}
