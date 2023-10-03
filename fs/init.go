package fs

import (
	"fmt"
	"os"
)

type Stat struct {
	path string
}

func (s *Stat) GetReadableSize() (string, error) {
	size, err := s.GetSize()
	if err != nil {
		return "", err
	}
	return HumanReadableSize(size), nil
}

func (s *Stat) GetSize() (int64, error) {
	fileInfo, err := os.Stat(s.path)
	if err != nil {
		return 0, err
	}
	return fileInfo.Size(), nil
}

func HumanReadableSize(sizeInput int64) string {
	size := float64(sizeInput)
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}
	base := 1024.0
	i := 0
	for size >= base && i < len(units)-1 {
		size = size / base
		i++
	}
	return fmt.Sprintf("%.2f%s", size, units[i])
}

func GetStat(path string) (*Stat, error) {
	// check if file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}
	return &Stat{path: path}, nil
}
