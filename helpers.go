package loadenv

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

const regex = `([\w]+=[\/\w]+)`

func testEnvString(s string) bool {
	r := regexp.MustCompile(regex)
	return r.MatchString(s)
}

func loadEnvFile(f string) (content string, err error) {
	bs, err := ioutil.ReadFile(f)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func getEnvLines(s string) []string {
	r := regexp.MustCompile(regex)
	return r.FindAllString(s, -1)
}

func getKeyValues(s string) (key string, value string, err error) {
	if ok := testEnvString(s); ok {
		kvp := strings.Split(s, "=")
		return kvp[0], kvp[1], nil
	}
	return "", "", errors.New(fmt.Sprintf("\"%v\" does not match test", s))
}
