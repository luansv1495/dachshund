package core

import (
	"errors"
	"fmt"
	"strings"
)

func ParseNodeID(nodeID string) (*NodePath, error) {
	parts := strings.Split(nodeID, "/")
	if len(parts) == 0 {
		return nil, errors.New("invalid nodeID")
	}

	path := &NodePath{}

	for _, p := range parts {
		if strings.HasPrefix(p, "conn:") {
			path.ConnectionID = strings.TrimPrefix(p, "conn:")
		} else if strings.HasPrefix(p, "database:") {
			path.Database = strings.TrimPrefix(p, "database:")
		} else if strings.HasPrefix(p, "schema:") {
			path.Schema = strings.TrimPrefix(p, "schema:")
		} else if strings.HasPrefix(p, "table:") {
			path.Table = strings.TrimPrefix(p, "table:")
		} else if strings.HasPrefix(p, "column:") {
			path.Column = strings.TrimPrefix(p, "column:")
		}
	}

	path.Path = nodeID
	return path, nil
}

func BuildDBPath(parentID, nodeType, name string) string {
	if parentID == "" {
		if nodeType == "database" {
			return fmt.Sprintf("conn:%s/db:%s", name, name) // root database node
		}
		return fmt.Sprintf("%s:%s", nodeType, name)
	}
	return fmt.Sprintf("%s/%s:%s", parentID, nodeType, name)
}
