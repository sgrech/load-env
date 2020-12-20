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

	if _, err := loadEnvFile("/bad/path"); err == nil {
		t.Fatal("Expected error in loading \"/bad/path\" but got none")
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

func TestGetEnvLines(t *testing.T) {
	if s, err := loadEnvFile(SLEF); err == nil {
		m := getEnvLines(s)
		if m != nil && len(m) != 1 {
			t.Fatalf("Expected slice of length 1 but got %d", len(m))
		}
		c := "TESTKEY=TESTVALUE"
		if m[0] != c {
			t.Fatalf("Error reading file contents, expected \"%v\", got  \"%v\"", c, m[0])
		}
	} else {
		t.Fatal(err)
	}
	if s, err := loadEnvFile(MLEF); err == nil {
		m := getEnvLines(s)
		if m != nil && len(m) != 3 {
			t.Fatalf("Expected slice of length 3 but got %d", len(m))
		}
		c1 := "TESTKEY1=TESTVALUE1"
		c2 := "TESTKEY2=TESTVALUE2"
		c3 := "TESTKEY3=TESTVALUE3"
		if m[0] != c1 {
			t.Fatalf("Error reading file contents, expected \"%v\", got  \"%v\"", c1, m[0])
		}
		if m[1] != c2 {
			t.Fatalf("Error reading file contents, expected \"%v\", got  \"%v\"", c2, m[1])
		}
		if m[2] != c3 {
			t.Fatalf("Error reading file contents, expected \"%v\", got  \"%v\"", c3, m[2])
		}
	} else {
		t.Fatal(err)
	}
	if s, err := loadEnvFile(PVEF); err == nil {
		m := getEnvLines(s)
		if m != nil && len(m) != 1 {
			t.Fatalf("Expected slice of length 1 but got %d", len(m))
		}
		c := "TESTKEY=/test/value"
		if m[0] != c {
			t.Fatalf("Error reading file contents, expected \"%v\", got  \"%v\"", c, m[0])
		}
	} else {
		t.Fatal(err)
	}
}

func TestGetKeyValues(t *testing.T) {
	if key, value, err := getKeyValues("TESTKEY=TESTVALUE"); err == nil {
		if key != "TESTKEY" {
			t.Fatalf("Expected key to be \"TESTKEY\" but got %v", key)
		}
		if value != "TESTVALUE" {
			t.Fatalf("Expected value to be \"TESTVALUE\" but got %v", value)
		}
	} else {
		t.Fatal(err)
	}
	if _, _, err := getKeyValues("TESTKEY-TESTVALUE"); err != nil {
		tErr := "\"TESTKEY-TESTVALUE\" does not match test"
		if err.Error() != tErr {
			t.Fatal("Unexpected error message")
		}
	} else {
		t.Fatal("Expected error in parsing \"TESTKEY-TESTVALUE\" but got none")
	}
}
