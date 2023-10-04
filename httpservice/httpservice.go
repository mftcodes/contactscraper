package httpservice

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func ClientSetup() (c http.Client) {
	return http.Client{Timeout: time.Duration(5) * time.Second}
}

func GetAction(uri string) (body []byte, err error) {
	c := ClientSetup()
	var fulluri string

	hasPrefix := strings.HasPrefix(uri, "https://")
	if hasPrefix {
		fulluri = uri
	} else {
		fulluri = fmt.Sprintf("%s%s", "https://", uri)
	}

	r, err := c.Get(fulluri)
	if err != nil {
		return body, err
	}
	defer r.Body.Close()
	body, err = ioutil.ReadAll(r.Body)

	if err != nil {
		return body, err
	}

	return body, nil
}
