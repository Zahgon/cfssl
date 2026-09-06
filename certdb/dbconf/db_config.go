package dbconf

import (
	"github.com/jmoiron/sqlx"
)

type DBConfig struct {
	DriverName     string `json:"driver"`
	DataSourceName string `json:"data_source"`
}

func LoadFile(path string) (cfg *DBConfig, err error) { _ = "STUB: not implemented"; return nil, nil }

func DBFromConfig(path string) (db *sqlx.DB, err error) { _ = "STUB: not implemented"; return nil, nil }
