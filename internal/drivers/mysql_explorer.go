package drivers

import (
	"context"
	"dachshund/internal/core"
)

type MySqlExplorer struct {
	driver core.DatabaseDriver
}

func NewMySqlExplorer(driver core.DatabaseDriver) core.Explorer {
	return &MySqlExplorer{
		driver: driver,
	}
}

func (d *MySqlExplorer) GetChildren(ctx context.Context, nodeID string) ([]core.TreeNode, error) {
	return nil, nil
}
