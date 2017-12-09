package model

type Config struct {
	CKey string `gorm:"column:ckey"`
	CSet string `gorm:"column:cset"`
}

func (Config) TableName() string {
	return "config"
}
