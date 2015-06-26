package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"io/ioutil"
)

func seedDatabase() {
	db, err := gorm.Open("postgres", "user=postgres dbname=ambition password=ambition")
	db.DB()

	seedJSON, err := ioutil.ReadFile("./config/seed.json")
	check(err)
	fmt.Print(string(seedJSON))

	var seed []Event

	json.Unmarshal(seedJSON, &seed)

	fmt.Print(seed)
}
