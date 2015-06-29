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
	db.d.CreateTable(&Action{})
	db.d.CreateTable(&Occurrence{})

}

func (db Database) clearTables() {
	db.d.Find(&Action{}).Delete(&Action{})
	db.d.Find(&Occurrence{}).Delete(&Occurrence{})
}

func (db Database) dropTables() {
	db.d.DropTable(&Action{})
	db.d.DropTable(&Occurrence{})
}

func (db Database) seedTables() {
	seedJSON, err := ioutil.ReadFile("./config/seed.json")
	check(err)
	fmt.Print(string(seedJSON))

	var seed []Action

	json.Unmarshal(seedJSON, &seed)

	fmt.Print(seed)

	for _, action := range seed {
		db.d.Create(&action)
	}
}
