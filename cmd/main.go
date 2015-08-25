package main

import (
	"log"
	"os"

	"github.com/adamryman/ambition"
)

var port = os.Getenv("ambition_port")

func main() {
	if len(os.Args) > 1 {
		if err := ambition.CallCommand(os.Args[1]); err != nil {
			log.Fatal(err)
		}
		return
	}
	ambition.Run()
}
