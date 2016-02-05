package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/headzoo/skeleton-go-react/app"
	"github.com/jinzhu/gorm"
)

var DB gorm.DB

func InitDatabase() error {
	var err error
	DB, err = gorm.Open(
		app.Settings.Database.Driver,
		app.Settings.Database.Dsn,
	)
	if err != nil {
		return err
	}
	return nil
}
