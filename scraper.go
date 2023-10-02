package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"scraper/helpers"
	"scraper/readservice"
)

func main() {
	isdebug := flag.Bool("d", false, "Sets program in debug mode")
	flag.Parse()

	readdir := helpers.SetReadDir()
	writedir := helpers.SetWriteDir()
	var urls []string
	var err error

	// DEBUG PRINTING REMOVE
	fmt.Printf("Read Dir: %s\n", readdir)
	fmt.Printf("Write Dir: %s\n", writedir)

	if *isdebug {
		if err = helpers.SetupDebugDirs(); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		readdir = helpers.SetDebugReadDir()
		writedir = helpers.SetDebugWriteDir()
	}

	files, err := readservice.ReadAllFiles(readdir)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	for _, f := range files {
		fileuri := fmt.Sprintf("%s%s", readdir, f)
		urls, err = readservice.ReadUrlsFromFile(fileuri)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		// DEBUG PRINTING REMOVE
		fmt.Printf("File: %s\n", f)
		for i, u := range urls {
			fmt.Printf("    Url %d: %s\n", i, u)
		}
	}
}
