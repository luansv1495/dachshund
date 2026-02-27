package core

import "context"

type DatabaseDriver interface {
	Connect(config ConnectionConfig) error
	Close() error
	Ping(ctx context.Context) error
	Query(ctx context.Context, sql string, args ...any) (Rows, error)
	Exec(ctx context.Context, sql string, args ...any) (Result, error)
	CurrentDatabase() string
	ConnectToDatabase(dbName string) error
}
