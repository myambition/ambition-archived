package main

import (
	"github.com/adamryman/ambition"
	"log"
	"os"
)

func main() {
	ambition.Init()
	if len(os.Args) > 1 {
		if err := ambition.CallCommand(os.Args[1]); err != nil {
			log.Fatal(err)
		}
		return
	}
	ambition.Run()
}
