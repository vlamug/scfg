package model

// Config is a config model
type Config struct {
	CKey string `gorm:"column:ckey"`
	CSet string `gorm:"column:cset"`
}

// TableName return table name
func (Config) TableName() string {
	return "config"
}
