package storage

import (
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"small_imgbed/config"
)

func Save(filename string, f multipart.File) string {
	b, err := io.ReadAll(f)
	if err != nil {
		log.Panicln(err)
	}
	f.Close()
	newFilePath := filepath.Join(config.C.Data.Path, filename)
	err = os.MkdirAll(filepath.Dir(newFilePath), os.ModePerm)
	if err != nil {
		log.Panicln(err)
	}
	err = os.WriteFile(newFilePath, b, os.ModePerm)
	if err != nil {
		log.Panicln(err)
	}
	return filename
}
func Get(p string) []byte {
	b, err := os.ReadFile(filepath.Join(config.C.Data.Path, p))
	if err != nil {
		log.Panicln(err)
	}
	return b
}
func FileExit(p string) bool {
	s, err := os.Stat(filepath.Join(config.C.Data.Path, p))
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		log.Panicln(err)
	}
	return !s.IsDir()
}
