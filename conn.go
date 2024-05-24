package main

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func conn() *gorm.DB {
	dsn := DBuser + ":" + DBpass + "@tcp(" + DatabaseIP + ":" + DatabasePort + ")/lnoblog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return db
}
