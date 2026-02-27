package drivers

import (
	"context"
	"dachshund/internal/core"
)

type SqliteExplorer struct {
	driver core.DatabaseDriver
}

func NewSqliteExplorer(driver core.DatabaseDriver) core.Explorer {
	return &SqliteExplorer{
		driver: driver,
	}
}

func (d *SqliteExplorer) GetChildren(ctx context.Context, nodeID string) ([]core.TreeNode, error) {
	return nil, nil
}
