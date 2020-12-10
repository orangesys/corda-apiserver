package filesystem

import (
	"os"
	"path/filepath"
)

//WriteToFile ...
func WriteToFile(filePath string, bytes []byte) error {
	path := filepath.Dir(filePath)
	if !exist(path) {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return err
		}
	}
	if !exist(filePath) {
		f, err := os.Create(filePath)
		defer f.Close()
		if err != nil {
			return err
		}
		if _, err = f.Write(bytes); err != nil {
			return err
		}
	} else {
		f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0600)
		defer f.Close()
		if err != nil {
			return err
		}
		if _, err = f.Write(bytes); err != nil {
			return err
		}
	}
	return nil
}

func exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
