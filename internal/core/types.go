package core

type Rows interface {
	Next() bool
	Scan(dest ...any) error
	Close()
	Err() error
}

type Result interface{}

type TreeNode struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"` // connection, database, schema, table, column
	Metadata    string `json:"metadata"`
	HasChildren bool   `json:"hasChildren"`
}

type NodePath struct {
	ConnectionID string
	Database     string
	Schema       string
	Table        string
	Column       string
	Path         string // caminho completo para gerar ID
}
