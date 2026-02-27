package connection

import (
	"context"
	"dachshund/internal/core"
	"dachshund/internal/factory"
	"errors"
	"sync"

	"github.com/google/uuid"
)

type ManagedConnection struct {
	Config   core.ConnectionConfig
	Driver   core.DatabaseDriver
	Explorer core.Explorer
}

type Manager struct {
	connections map[string]*ManagedConnection
	mu          sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{
		connections: make(map[string]*ManagedConnection),
	}
}

func (m *Manager) Create(config core.ConnectionConfig) (*core.ConnectionConfig, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	config.ID = uuid.NewString()

	if _, exists := m.connections[config.ID]; exists {
		return nil, errors.New("connection already exists")
	}

	driver, err := factory.NewDriver(config)
	if err != nil {
		return nil, err
	}

	if err := driver.Connect(config); err != nil {
		return nil, err
	}

	explorer, err := factory.NewExplorer(config, driver)
	if err != nil {
		return nil, err
	}

	m.connections[config.ID] = &ManagedConnection{
		Config:   config,
		Driver:   driver,
		Explorer: explorer,
	}

	return &config, nil
}

func (m *Manager) Get(id string) (core.DatabaseDriver, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	conn, ok := m.connections[id]
	if !ok {
		return nil, errors.New("connection not found")
	}

	return conn.Driver, nil
}

func (m *Manager) ListConnections() []core.ConnectionConfig {
	m.mu.RLock()
	defer m.mu.RUnlock()

	list := make([]core.ConnectionConfig, 0)

	for _, conn := range m.connections {
		list = append(list, conn.Config)
	}

	return list
}

func (m *Manager) Close(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	conn, ok := m.connections[id]
	if !ok {
		return errors.New("connection not found")
	}

	conn.Driver.Close()
	delete(m.connections, id)
	return nil
}

func (m *Manager) Ping(id string) error {
	conn, err := m.Get(id)
	if err != nil {
		return err
	}

	return conn.Ping(context.Background())
}

func (m *Manager) Test(config core.ConnectionConfig) error {
	driver, err := factory.NewDriver(config)
	if err != nil {
		return err
	}

	// tenta conectar
	if err := driver.Connect(config); err != nil {
		return err
	}

	// tenta ping
	if err := driver.Ping(context.Background()); err != nil {
		driver.Close()
		return err
	}

	// fecha depois de testar
	driver.Close()

	return nil
}

func (m *Manager) GetExplorer(id string) (core.Explorer, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	conn, ok := m.connections[id]
	if !ok {
		return nil, errors.New("connection not found")
	}

	return conn.Explorer, nil
}
