package files

import (
	"fmt"
	"os"

	"github.com/BomScoob12/homelab-file-manager/internal/fs"
)

type FileService struct{}

func (s *FileService) List() ([]os.DirEntry, error) {
	return fs.ListFile("/")
}

func (s *FileService) ReadFile() {
	fmt.Println("READ file: [1]")
}

func (s *FileService) DeleteFile(target string) {
	fmt.Printf("DELETE file: %s", target)
}
