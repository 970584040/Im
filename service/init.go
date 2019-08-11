package service

import (
	"../model"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

var DbEngin *xorm.Engine

func init() {
	DriveName := "mysql"
	DsName := "root:root@(localhost:3306)/im?charset=utf8"
	err := errors.New("")
	DbEngin, err = xorm.NewEngine(DriveName, DsName)
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	DbEngin.ShowSQL(true)

	DbEngin.Sync2(new(model.User), new(model.Contact), new(model.Community))
}
