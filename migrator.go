package hive

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/migrator"
)

type Migrator struct {
	migrator.Migrator
	*Dialector
}

func (m Migrator) CurrentDatabase() (name string) {
	m.handleError(m.DB.Raw("SELECT current_database()").Row().Scan(&name))
	return
}

func (m Migrator) HasTable(value interface{}) bool {
	var name string
	m.handleError(m.RunWithValue(value, func(stmt *gorm.Statement) error {
		//currentDatabase := m.DB.Migrator().CurrentDatabase()
		return m.DB.Raw(
			// TODO: support args
			fmt.Sprintf("SHOW TABLES LIKE '%s'", stmt.Table),
		).Row().Scan(&name)
	}))
	return name != ""
}
