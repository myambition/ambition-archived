package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func clearDatabase() {
	db, err := gorm.Open("postgres", "user=postgres dbname=ambition password=ambition")
	check(err)
	db.DB()

	db.DropTable(&Event{})
	db.DropTable(&EventTime{})

}
