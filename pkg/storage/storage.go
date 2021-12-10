package storage

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"io"
	"strings"
)

type UploadInput struct {
	File        io.Reader
	Name        string
	Size        int64
	ContentType string
}

type Provider interface {
	Upload(ctx context.Context, input UploadInput) (string, error)
}

func (input *UploadInput) MakeRandomName() (string, error) {
	newName := genUUID()

	splitParts := strings.Split(input.Name, ".")

	if len(splitParts) > 1 {
		fileExt := splitParts[len(splitParts)-1]

		newName = fmt.Sprintf("%s.%s", newName, fileExt)
	}

	return newName, nil
}

func genUUID() string {
	uuidWithHyphen := uuid.New()
	fmt.Println(uuidWithHyphen)
	guid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)

	return guid
}
