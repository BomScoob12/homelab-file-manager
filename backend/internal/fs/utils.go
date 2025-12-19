package fs

import (
	"fmt"
	"os"
)

func ReadFile(path string) {
	fmt.Println(path)
}

func ListFile(path string) ([]os.DirEntry, error) {
	return os.ReadDir("/WorkDir")
}
