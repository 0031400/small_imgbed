package storage

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"small_imgbed/config"
)

func Save(filename string, f multipart.File) (string, error) {
	b, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}
	f.Close()
	newFilePath := filepath.Join(config.C.Data.Path, filename)
	err = os.MkdirAll(filepath.Dir(newFilePath), os.ModePerm)
	if err != nil {
		return "", err
	}
	err = os.WriteFile(newFilePath, b, os.ModePerm)
	if err != nil {
		return "", err
	}
	return filename, nil
}
func GetPath(p string) string {
	return filepath.Join(config.C.Data.Path, p)

}
func Get(p string) ([]byte, error) {
	b, err := os.ReadFile(p)
	if err != nil {
		return nil, err
	}
	return b, nil
}
func FileExit(p string) (bool, error) {
	s, err := os.Stat(p)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return !s.IsDir(), nil
}
