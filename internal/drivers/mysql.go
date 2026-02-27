package drivers

import (
	"context"
	"dachshund/internal/core"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type MySqlDriver struct {
	conn *sql.DB
}

func (d *MySqlDriver) Connect(config core.ConnectionConfig) error {
	return nil
}

func (d *MySqlDriver) Close() error {
	return nil
}

func (d *MySqlDriver) Ping(ctx context.Context) error {
	return nil
}

func (d *MySqlDriver) Query(ctx context.Context, sql string, args ...any) (core.Rows, error) {
	return nil, nil
}

func (d *MySqlDriver) Exec(ctx context.Context, sql string, args ...any) (core.Result, error) {
	return nil, nil
}

func (d *MySqlDriver) ConnectToDatabase(dbName string) error {
	return nil
}

func (d *MySqlDriver) CurrentDatabase() string {
	return ""
}
