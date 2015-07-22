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
	db.d.CreateTable(&User{})
	db.d.CreateTable(&Set{})
	db.d.CreateTable(&Action{})
	db.d.CreateTable(&Occurrence{})

}

func (db Database) dropTables() {
	db.d.DropTable(&User{})
	db.d.DropTable(&Set{})
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

	db.d.Create(&Set{SetName: "Health", Actions: seed})
}

func (db Database) refreshTables() {
	database.dropTables()
	database.createTables()
	database.seedTables()
}
