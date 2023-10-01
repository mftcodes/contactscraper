package main

import (
	"flag"
	"fmt"
	"log"

	"scraper/helpers"
)

func main() {
	isdebug := flag.Bool("d", false, "Sets program in debug mode")
	flag.Parse()

	if *isdebug {
		if err := helpers.SetupDebugDirs(); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Look, Ma, No hands!")
}
