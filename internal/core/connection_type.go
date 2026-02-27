package core

type ConnectionType struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	DefaultPort int    `json:"defaultPort"`
}
