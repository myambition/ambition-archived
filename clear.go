package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func clearTables() {
	db, err := gorm.Open("postgres", "user=postgres dbname=ambition password=ambition")
	check(err)
	db.DB()

	db.Find(&Event{}).Delete(&Event{})
	db.Find(&EventTime{}).Delete(&EventTime{})
}
