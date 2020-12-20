package loadenv

import (
	"os"
)

// Read contents of file f and parse contents to set respective environment
// variables. File contents should be in the format KEY=VALUE, each separated
// by a newline.  An error is returned if the file does not exist, or the
// contents do not followed the right formatting.
func LoadEnvVars(f string) error {
	if c, err := loadEnvFile(f); err == nil {
		lns := getEnvLines(c)
		for _, ln := range lns {
			if key, val, err := getKeyValues(ln); err == nil {
				if err := os.Setenv(key, val); err != nil {
					return err
				}
			} else {
				return err
			}
		}
	} else {
		return err
	}
	return nil
}
