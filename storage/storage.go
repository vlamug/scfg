package storage

import (
	"github.com/jinzhu/gorm"
	"github.com/vlamug/scfg/model"
)

type Storage interface {
	Get(ckey string) model.Config
}

type PostgresStorage struct {
	db *gorm.DB
}

func NewPostgresStorage(db *gorm.DB) *PostgresStorage {
	return &PostgresStorage{db: db}
}

func (st *PostgresStorage) Get(ckey string) model.Config {
	var cfg model.Config
	st.db.First(&cfg, "ckey = ?", ckey)

	return cfg
}
