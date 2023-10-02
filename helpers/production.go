package helpers

import (
	"os"
)

func SetupDirs() error {
	readdir := SetReadDir()
	err := os.MkdirAll(readdir, 0750)
	if err != nil && !os.IsExist(err) {
		return err
	}

	writedir := SetWriteDir()
	err = os.MkdirAll(writedir, 0750)
	if err != nil && !os.IsExist(err) {
		return err
	}
	return nil
}

func SetReadDir() string {
	return "/srv/scrape.in/"
}

func SetWriteDir() string {
	return "/srv/scrape.out/"
}
