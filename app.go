package main

import (
	"context"
	"dachshund/internal/connection"
	"dachshund/internal/core"
	"dachshund/internal/factory"
	"fmt"
)

// App struct
type App struct {
	ctx     context.Context
	manager *connection.Manager
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		manager: connection.NewManager(),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) CreateConnection(config core.ConnectionConfig) (*core.ConnectionConfig, error) {
	return a.manager.Create(config)
}

func (a *App) ListConnections() []core.ConnectionConfig {
	return a.manager.ListConnections()
}

func (a *App) ListConnectionTypes() []core.ConnectionType {
	return factory.SupportedConnectionTypes()
}

func (a *App) TestConnection(config core.ConnectionConfig) error {
	return a.manager.Test(config)
}

func (a *App) GetChildren(nodeID string) ([]core.TreeNode, error) {
	path, err := core.ParseNodeID(nodeID)

	if err != nil {
		return nil, fmt.Errorf("invalid nodeID: missing connection")
	}

	explorer, err := a.manager.GetExplorer(path.ConnectionID)
	if err != nil {
		return nil, err
	}

	return explorer.GetChildren(a.ctx, nodeID)
}
