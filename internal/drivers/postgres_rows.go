package drivers

import "github.com/jackc/pgx/v5"

type PostgresRows struct {
	rows pgx.Rows
}

func (r *PostgresRows) Next() bool {
	return r.rows.Next()
}

func (r *PostgresRows) Scan(dest ...any) error {
	return r.rows.Scan(dest...)
}

func (r *PostgresRows) Close() {
	r.rows.Close()
}

func (r *PostgresRows) Err() error {
	return r.rows.Err()
}
