package factory

import (
	"dachshund/internal/core"
	"dachshund/internal/drivers"
	"errors"
)

func NewExplorer(config core.ConnectionConfig, driver core.DatabaseDriver) (core.Explorer, error) {
	switch config.Type {
	case "postgres":
		return drivers.NewPostgresExplorer(driver), nil
	case "mysql":
		return drivers.NewMySqlExplorer(driver), nil
	case "sqlite":
		return drivers.NewSqliteExplorer(driver), nil
	default:
		return nil, errors.New("unsupported driver explorer")
	}
}
