package files

import "fmt"

type FileService struct{}

func (s *FileService) List() {
	fmt.Println("LIST of file: [1,2,3]")
}

func (s *FileService) ReadFile() {
	fmt.Println("READ file: [1]")
}

func (s *FileService) DeleteFile() {
	fmt.Println("DELETE file: [1,2,3]")
}
