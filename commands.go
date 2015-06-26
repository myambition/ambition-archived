package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"io/ioutil"
	"reflect"
)

type Database struct {
	d gorm.DB
}

func (db Database) createTables() {
	fmt.Print(reflect.TypeOf(db))
	db.d.CreateTable(&Event{})
	db.d.CreateTable(&EventTime{})

}

func (db Database) clearTables() {
	db.d.Find(&Event{}).Delete(&Event{})
	db.d.Find(&EventTime{}).Delete(&EventTime{})
}

func (db Database) dropTables() {
	db.d.DropTable(&Event{})
	db.d.DropTable(&EventTime{})
}

func (db Database) seedTables() {
	seedJSON, err := ioutil.ReadFile("./config/seed.json")
	check(err)
	fmt.Print(string(seedJSON))

	var seed []Event

	json.Unmarshal(seedJSON, &seed)

	fmt.Print(seed)

	for _, event := range seed {
		db.d.Create(&event)
	}
}
