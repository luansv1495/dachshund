package factory

import (
	"dachshund/internal/core"
	"dachshund/internal/drivers"
	"errors"
)

func NewDriver(config core.ConnectionConfig) (core.DatabaseDriver, error) {
	switch config.Type {
	case "postgres":
		return &drivers.PostgresDriver{}, nil
	case "mysql":
		return &drivers.MySqlDriver{}, nil
	case "sqlite":
		return &drivers.SqliteDriver{}, nil
	default:
		return nil, errors.New("unsupported driver")
	}
}

func SupportedConnectionTypes() []core.ConnectionType {
	return []core.ConnectionType{
		{
			ID:          "postgres",
			Name:        "PostgreSQL",
			DefaultPort: 5432,
		},
		{
			ID:          "mysql",
			Name:        "MySQL",
			DefaultPort: 3306,
		},
		{
			ID:          "sqlite",
			Name:        "SQLite",
			DefaultPort: 0,
		},
	}
}
