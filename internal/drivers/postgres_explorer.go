package drivers

import (
	"context"
	"dachshund/internal/core"
	"errors"
)

type PostgresExplorer struct {
	driver core.DatabaseDriver
}

func NewPostgresExplorer(driver core.DatabaseDriver) core.Explorer {
	return &PostgresExplorer{
		driver: driver,
	}
}

func (e *PostgresExplorer) GetChildren(ctx context.Context, nodeID string) ([]core.TreeNode, error) {
	path, err := core.ParseNodeID(nodeID)
	if err != nil {
		return nil, err
	}

	// Se não tem database, retorna databases
	if path.Database == "" {
		return e.listDatabases(ctx, path)
	}

	// Se não tem schema, retorna schemas do database
	if path.Schema == "" {
		return e.listSchemas(ctx, path)
	}

	// Se não tem table, retorna tables do schema
	if path.Table == "" {
		return e.listTables(ctx, path)
	}

	// Se chegou aqui, retorna columns da tabela
	return e.listColumns(ctx, path)
}

func (e *PostgresExplorer) listDatabases(ctx context.Context, path *core.NodePath) ([]core.TreeNode, error) {
	rows, err := e.driver.Query(ctx, `
		SELECT datname
		FROM pg_database
		WHERE datistemplate = false
		ORDER BY datname
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var nodes []core.TreeNode

	for rows.Next() {
		var dbName string
		if err := rows.Scan(&dbName); err != nil {
			return nil, err
		}

		id := core.BuildDBPath(path.Path, "database", dbName)
		nodes = append(nodes, core.TreeNode{
			ID:          id,
			Name:        dbName,
			Type:        "database",
			HasChildren: true,
		})
	}

	return nodes, nil
}

func (e *PostgresExplorer) listSchemas(ctx context.Context, path *core.NodePath) ([]core.TreeNode, error) {
	if path.Database == "" {
		return nil, errors.New("database not set")
	}

	// troca de database se necessário
	if err := e.driver.ConnectToDatabase(path.Database); err != nil {
		return nil, err
	}

	rows, err := e.driver.Query(ctx, `
		SELECT schema_name
		FROM information_schema.schemata
		ORDER BY schema_name
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var nodes []core.TreeNode

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}

		id := core.BuildDBPath(path.Path, "schema", name)
		nodes = append(nodes, core.TreeNode{
			ID:          id,
			Name:        name,
			Type:        "schema",
			HasChildren: true,
		})
	}

	return nodes, nil
}

func (e *PostgresExplorer) listTables(ctx context.Context, path *core.NodePath) ([]core.TreeNode, error) {
	if path.Database == "" || path.Schema == "" {
		return nil, errors.New("database or schema not set")
	}

	rows, err := e.driver.Query(ctx, `
		SELECT table_name
		FROM information_schema.tables
		WHERE table_schema = $1
		  AND table_type = 'BASE TABLE'
		ORDER BY table_name
	`, path.Schema)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var nodes []core.TreeNode

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}

		id := core.BuildDBPath(path.Path, "table", name)
		nodes = append(nodes, core.TreeNode{
			ID:          id,
			Name:        name,
			Type:        "table",
			HasChildren: true,
		})
	}

	return nodes, nil
}

func (e *PostgresExplorer) listColumns(ctx context.Context, path *core.NodePath) ([]core.TreeNode, error) {
	if path.Database == "" || path.Schema == "" || path.Table == "" {
		return nil, errors.New("database/schema/table not set")
	}

	rows, err := e.driver.Query(ctx, `
		SELECT column_name, udt_name
		FROM information_schema.columns
		WHERE table_schema = $1
		  AND table_name = $2
		ORDER BY ordinal_position
	`, path.Schema, path.Table)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var nodes []core.TreeNode

	for rows.Next() {
		var name string
		var datatype string
		if err := rows.Scan(&name, &datatype); err != nil {
			return nil, err
		}

		id := core.BuildDBPath(path.Path, "column", name)
		nodes = append(nodes, core.TreeNode{
			ID:          id,
			Name:        name,
			Metadata:    datatype,
			Type:        "column",
			HasChildren: false,
		})
	}

	return nodes, nil
}
