package util

import (
	"os"
)

func WriteFile(f string, data []byte) error {
	return os.WriteFile(f, data, 0666)
}
