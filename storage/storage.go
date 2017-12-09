package storage

import (
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/vlamug/scfg/model"
	"github.com/vlamug/scfg/request"
)

type Storage interface {
	Get(request request.GetRequest) (model.Config)
}

type PostgresStorage struct {
	db *gorm.DB
}

func NewPostgresStorage(db *gorm.DB) *PostgresStorage {
	return &PostgresStorage{db: db}
}

func (st *PostgresStorage) Get(req request.GetRequest) (model.Config) {
	var cfg model.Config
	st.db.First(&cfg, "ckey = ?", strings.Join([]string{req.Type, req.Data}, ":"))

	return cfg
}