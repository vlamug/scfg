package config

import (
	"io/ioutil"
	"encoding/json"
)

// Database is config of connection to database
type Database struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"db_name"`
}

// LoadDatabaseConfig loads database config
func LoadDatabaseConfig(path string) (*Database, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg := &Database{}
	json.Unmarshal(content, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
