package model

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Db *gorm.DB

func NewMySqlConn() (err error) {
	Db, err = gorm.Open(mysql.Open(viper.GetString("DB.Url")), &gorm.Config{})
	if err != nil {
		println("failed start app")
		return err
	}
	if sqlDb, err := Db.DB(); err == nil {
		sqlDb.SetMaxIdleConns(viper.GetInt("DB.MaxIdleConns"))
		sqlDb.SetMaxOpenConns(viper.GetInt("DB.MaxOpenConns"))
	} else {
		log.Fatal(err)
		return err
	}

	return nil
}

func DB() *gorm.DB {
	return Db
}
