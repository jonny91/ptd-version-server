package controlls

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	DB *gorm.DB
)

func InitDB() error {
	DB, err = gorm.Open("mysql", "root:mailingjoy.123@(39.105.34.250:13306)/ptd?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return err
	}
	return nil
}
