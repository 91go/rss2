package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("./rss2.db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err == nil {
		DB = db
		err = db.AutoMigrate(&Yearly{}, &Everyday{})
		if err != nil {
			return nil, err
		}
		return DB, err
	}

	return nil, err
}
