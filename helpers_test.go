package loadenv

import (
	"testing"
)

const SLEF = "./test/data/single_line.env"
const MLEF = "./test/data/multi_line.env"
const PVEF = "./test/data/path_val.env"

func TestLoadEnvFile(t *testing.T) {
	if s, err := loadEnvFile(SLEF); err == nil {
		c := "TESTKEY=TESTVALUE\n"
		if s != c {
			t.Fatalf("Error reading file contents, expected \"%v\", got  \"%v\"", c, s)
		}
	} else {
		t.Fatal(err)
	}
	if s, err := loadEnvFile(MLEF); err == nil {
		c := "TESTKEY1=TESTVALUE1\nTESTKEY2=TESTVALUE2\nTESTKEY3=TESTVALUE3\n"
		if s != c {
			t.Fatalf("Error reading file contents, expected \"%v\", got  \"%v\"", c, s)
		}
	} else {
		t.Fatal(err)
	}
	if s, err := loadEnvFile(PVEF); err == nil {
		c := "TESTKEY=/test/value\n"
		if s != c {
			t.Fatalf("Error reading file contents, expected \"%v\", got  \"%v\"", c, s)
		}
	} else {
		t.Fatal(err)
	}
}

func TestEnvString(t *testing.T) {
	if s, err := loadEnvFile(SLEF); err == nil {
		m := testEnvString(s)
		if !m {
			t.Fatalf("Expected \"%v\" to pass regexp test but got %t", SLEF, m)
		}
	} else {
		t.Fatal(err)
	}
	if s, err := loadEnvFile(MLEF); err == nil {
		m := testEnvString(s)
		if !m {
			t.Fatalf("Expected \"%v\" to pass regexp test but got %t", MLEF, m)
		}
	} else {
		t.Fatal(err)
	}
	if s, err := loadEnvFile(PVEF); err == nil {
		m := testEnvString(s)
		if !m {
			t.Fatalf("Expected \"%v\" to pass regexp test but got %t", MLEF, m)
		}
	} else {
		t.Fatal(err)
	}
}
