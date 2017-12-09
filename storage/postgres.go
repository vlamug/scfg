package storage

import (
	"github.com/jinzhu/gorm"
	"github.com/vlamug/scfg/model"
)

// PostgresStorage is postgresql storage implementation
type PostgresStorage struct {
	db *gorm.DB
}

// NewPostgresStorage creates new postgresql storage implementation
func NewPostgresStorage(db *gorm.DB) *PostgresStorage {
	return &PostgresStorage{db: db}
}

// Get loads config from database
func (st *PostgresStorage) Get(ckey string) model.Config {
	var cfg model.Config
	st.db.First(&cfg, "ckey = ?", ckey)

	return cfg
}
