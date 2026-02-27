package drivers

import (
	"context"
	"dachshund/internal/core"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresDriver struct {
	pools  map[string]*pgxpool.Pool
	mu     sync.RWMutex
	config *core.ConnectionConfig
}

func (d *PostgresDriver) Connect(config core.ConnectionConfig) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.config = &config
	d.pools = make(map[string]*pgxpool.Pool)

	pool, err := d.createPool(config.Database)
	if err != nil {
		return err
	}

	d.pools[config.Database] = pool

	return nil
}

func (d *PostgresDriver) Close() error {
	pool, err := d.getPool(d.config.Database)
	if err != nil {
		return err
	}

	pool.Close()
	return nil
}

func (d *PostgresDriver) Ping(ctx context.Context) error {
	pool, err := d.getPool(d.config.Database)
	if err != nil {
		return err
	}

	return pool.Ping(ctx)
}

func (d *PostgresDriver) Query(ctx context.Context, sql string, args ...any) (core.Rows, error) {
	pool, err := d.getPool(d.config.Database)
	if err != nil {
		return nil, err
	}

	rows, err := pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	return &PostgresRows{rows: rows}, nil
}

func (d *PostgresDriver) Exec(ctx context.Context, sql string, args ...any) (core.Result, error) {
	pool, err := d.getPool(d.config.Database)
	if err != nil {
		return nil, err
	}

	cmd, err := pool.Exec(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

func (d *PostgresDriver) CurrentDatabase() string {
	if d.config == nil {
		return ""
	}
	return d.config.Database
}

func (d *PostgresDriver) ConnectToDatabase(dbName string) error {
	_, err := d.getPool(dbName)
	if err != nil {
		return err
	}

	d.mu.Lock()
	d.config.Database = dbName
	d.mu.Unlock()

	return err
}

func (d *PostgresDriver) createPool(dbName string) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		d.config.User, d.config.Password, d.config.Host, d.config.Port, dbName, d.config.SSLMode,
	)
	return pgxpool.New(context.Background(), dsn)
}

func (d *PostgresDriver) getPool(dbName string) (*pgxpool.Pool, error) {
	if dbName == "" {
		dbName = d.config.Database
	}

	d.mu.RLock()
	pool, ok := d.pools[dbName]
	d.mu.RUnlock()

	if ok {
		return pool, nil
	}

	newPool, err := d.createPool(dbName)
	if err != nil {
		return nil, err
	}
	d.pools[dbName] = newPool
	return newPool, nil
}
