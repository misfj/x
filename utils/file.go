package utils

import (
	"bytes"
	"io"
	"os"
)

func CreateAppendFile(path string, content []byte) error {
	writer, err := os.Create(path)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, bytes.NewReader(content))
	if err != nil {
		return err
	}
	return nil
}
