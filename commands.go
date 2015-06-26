package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"io/ioutil"
)

func createTables() {
	db, err := gorm.Open("postgres", "user=postgres dbname=ambition password=ambition")
	check(err)
	db.DB()

	db.CreateTable(&Event{})
	db.CreateTable(&EventTime{})

}

func clearTables() {
	db, err := gorm.Open("postgres", "user=postgres dbname=ambition password=ambition")
	check(err)
	db.DB()

	db.Find(&Event{}).Delete(&Event{})
	db.Find(&EventTime{}).Delete(&EventTime{})
}

func dropTables() {
	db, err := gorm.Open("postgres", "user=postgres dbname=ambition password=ambition")
	check(err)
	db.DB()

	db.DropTable(&Event{})
	db.DropTable(&EventTime{})
}

func seedTables() {
	db, err := gorm.Open("postgres", "user=postgres dbname=ambition password=ambition")
	db.DB()

	seedJSON, err := ioutil.ReadFile("./config/seed.json")
	check(err)
	fmt.Print(string(seedJSON))

	var seed []Event

	json.Unmarshal(seedJSON, &seed)

	fmt.Print(seed)

	for _, event := range seed {
		db.Create(&event)
	}
}
