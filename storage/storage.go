package storage

import (
	"github.com/vlamug/scfg/model"
)

// Storage is an interface for storage
type Storage interface {
	Get(ckey string) model.Config
}
