package shared

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // Ejemplo usando SQLite
)

func NewDB(dataSourceName string) (*sql.DB, error) {
	return sql.Open("sqlite3", dataSourceName)
}
