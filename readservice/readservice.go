package readservice

import (
	"bufio"
	"fmt"
	"os"
)

func ReadAllFiles(diruri string) (filenames []string, err error) {
	files, err := os.ReadDir(diruri)
	if err != nil {
		return filenames, err
	}

	for _, f := range files {
		filenames = append(filenames, fmt.Sprintln(f.Name()))
	}

	return filenames, nil
}

func ReadUrlsFromFile(fileuri string) (urls []string, err error) {
	file, err := os.Open(fileuri)
	if err != nil {
		return urls, err
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		urls = append(urls, sc.Text())
	}
	if err = sc.Err(); err != nil {
		return urls, err
	}

	return urls, nil
}
