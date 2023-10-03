package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"scraper/helpers"
	"scraper/httpservice"
	"scraper/readservice"
	"scraper/writeservice"
)

func main() {
	isdebug := flag.Bool("d", false, "Sets program in debug mode")
	flag.Parse()

	readdir := helpers.SetReadDir()
	writedir := helpers.SetWriteDir()
	var urls []string
	var err error

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
		urls, err = readservice.ReadUrlsFromFile(strings.TrimSpace(fileuri))
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		// DEBUG PRINTING REMOVE
		fmt.Printf("File: %s\n", f)
		for i, u := range urls {
			fmt.Printf("    Url %d: %s\n", i, u)
			body, err := httpservice.GetAction(u)
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
			fileuri := fmt.Sprintf("%s%s", writedir, u)
			if err = writeservice.BodyToFile(body, fileuri, *isdebug); err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
		}
	}

}
