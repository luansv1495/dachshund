package drivers

import (
	"context"
	"dachshund/internal/core"
	"database/sql"

	_ "modernc.org/sqlite"
)

type SqliteDriver struct {
	conn *sql.DB
}

func (d *SqliteDriver) Connect(config core.ConnectionConfig) error {
	return nil
}

func (d *SqliteDriver) Close() error {
	return nil
}

func (d *SqliteDriver) Ping(ctx context.Context) error {
	return nil
}

func (d *SqliteDriver) Query(ctx context.Context, sql string, args ...any) (core.Rows, error) {
	return nil, nil
}

func (d *SqliteDriver) Exec(ctx context.Context, sql string, args ...any) (core.Result, error) {
	return nil, nil
}

func (d *SqliteDriver) ConnectToDatabase(dbName string) error {
	return nil
}

func (d *SqliteDriver) CurrentDatabase() string {
	return ""
}
