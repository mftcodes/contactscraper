package helpers

import (
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
	return nil
}

func SetDebugReadDir() string {
	return "/tmp/scrape.in/"
}

func SetDebugWriteDir() string {
	return "/tmp/scrape.out/"
}
