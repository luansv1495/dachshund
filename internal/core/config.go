package core

type ConnectionConfig struct {
	ID       string
	Type     string
	Host     string
	Port     int
	User     string
	Password string
	Database string
	SSLMode  string
}
