package storage

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"
)

type FileStorage struct {
	Bucket string
}

func NewFileStorage(bucket string) *FileStorage {
	return &FileStorage{
		Bucket: bucket,
	}
}

func (fs *FileStorage) Upload(ctx context.Context, input UploadInput) (string, error) {
	path := makeSubDirectoriesPath()
	fullPath := fs.Bucket + "/" + path

	if err := os.MkdirAll(fullPath, 0777); err != nil {
		return "", err
	}

	fileName, err := input.MakeRandomName()
	if err != nil {
		return "", err
	}

	fo, err := os.Create(fullPath + "/" + fileName)
	if err != nil {
		return "", err
	}
	defer func(fo *os.File) {
		err := fo.Close()
		if err != nil {
			return
		}
	}(fo)

	return fmt.Sprintf("/%s/%s", path, fileName), nil
}

func makeSubDirectoriesPath() string {
	var path string

	now := time.Now()
	currentYear, currentMonth, currentDay := now.Date()

	path = fmt.Sprintf("%s/%s_%s",
		strconv.Itoa(currentYear),
		strconv.Itoa(int(currentMonth)),
		strconv.Itoa(currentDay))

	return path
}
