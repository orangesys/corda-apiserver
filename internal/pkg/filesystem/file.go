package filesystem

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
)

//WriteToFile ...
func WriteToFile(filePath string, bytes []byte) error {
	path := filepath.Dir(filePath)
	if !Exist(path) {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return err
		}
	}
	if !Exist(filePath) {
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

//Exist ...
func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

//CreateDir ...
func CreateDir(path string) error {
	if !Exist(path) {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

//Zip ..
func Zip(dst, src string) (err error) {
	fw, err := os.Create(dst)
	defer fw.Close()
	if err != nil {
		return err
	}

	zw := zip.NewWriter(fw)
	defer func() {
		if err := zw.Close(); err != nil {
			log.Println(err)
		}
	}()

	return filepath.Walk(src, func(path string, fi os.FileInfo, errBack error) (err error) {
		if errBack != nil {
			return errBack
		}

		fh, err := zip.FileInfoHeader(fi)
		if err != nil {
			return
		}

		fh.Name = filepath.Base(path)

		if fi.IsDir() {
			return
		}

		w, err := zw.CreateHeader(fh)
		if err != nil {
			return
		}

		if !fh.Mode().IsRegular() {
			return nil
		}

		fr, err := os.Open(path)
		defer fr.Close()
		if err != nil {
			return
		}

		if _, err = io.Copy(w, fr); err != nil {
			return err
		}

		return nil
	})
}
