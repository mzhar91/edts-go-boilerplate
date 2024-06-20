package credential

import (
	"fmt"
	
	"github.com/huandu/go-sqlbuilder"
)

func (m *psqlRepository) selectCondition(sql *sqlbuilder.SelectBuilder, value string, tipe string) {
	sql.Where(
		sql.Equal(fmt.Sprintf("c.%v", tipe), value),
	)
}
