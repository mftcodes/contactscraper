package helpers

import (
	"fmt"
	"os"
)

func SetupDebugDirs() error {
	readdir := SetDebugReadDir()
	err := os.MkdirAll(readdir, 0750)
	if err != nil && !os.IsExist(err) {
		return err
	}

	writedir := SetDebugWriteDir()
	err = os.MkdirAll(writedir, 0750)
	if err != nil && !os.IsExist(err) {
		return err
	}

	testfile := fmt.Sprintln("ubuntu.com\nwww.debian.org\nsystem76.com\n")

	fileuri := fmt.Sprintf("%s%s", readdir, "urls.txt")
	err = os.WriteFile(fileuri, []byte(testfile), 0660)
	if err != nil {
		return err
	}

	return nil
}

func SetDebugReadDir() string {
	return "/tmp/scrape.in/"
}

func SetDebugWriteDir() string {
	return "/tmp/scrape.out/"
}
