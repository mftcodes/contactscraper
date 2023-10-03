package writeservice

import (
	"os"
)

func BodyToFile(body []byte, fileuri string, isdebug bool) error {
	err := os.WriteFile(fileuri, []byte(body), 0660)
	if err != nil {
		return err
	}
	return nil
}
