package loadenv

import (
	"io/ioutil"
	"regexp"
)

func testEnvString(s string) bool {
	r := regexp.MustCompile(`([\w]+=[\/\w]+)`)
	return r.MatchString(s)
}

func loadEnvFile(f string) (string, error) {
	bs, err := ioutil.ReadFile(f)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}
