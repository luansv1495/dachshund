package core

import "context"

type Explorer interface {
	GetChildren(ctx context.Context, nodeID string) ([]TreeNode, error)
}
