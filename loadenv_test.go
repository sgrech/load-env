package loadenv

import (
	"os"
	"testing"
)

func TestLoadEnvVars(t *testing.T) {
	if err := LoadEnvVars(SLEF); err == nil {
		if val := os.Getenv("TESTKEY"); val != "TESTVALUE" {
			t.Fatalf("Error reading env variable contents, expected \"TESTVAL\", got  \"%v\"", val)
			t.Fatal()
		}
	} else {
		t.Fatal(err)
	}
}
