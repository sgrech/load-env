package loadenv

import (
	"io/ioutil"
	"regexp"
)

const regex = `([\w]+=[\/\w]+)`

func testEnvString(s string) bool {
	r := regexp.MustCompile(regex)
	return r.MatchString(s)
}

func loadEnvFile(f string) (string, error) {
	bs, err := ioutil.ReadFile(f)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func getEnvLines(f string) []string {
	r := regexp.MustCompile(regex)
	return r.FindAllString(f, -1)
}
