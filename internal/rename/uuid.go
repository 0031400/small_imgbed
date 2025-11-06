package rename

import (
	"crypto/rand"
	"fmt"
	"log"
	"path/filepath"
	"time"
)

func RenamePath(filename string) string {
	now := time.Now().UTC()
	return fmt.Sprintf("%02d/%02d/%02d/%s%s", now.Year(), now.Month(), now.Day(), UUID(), filepath.Ext(filename))
}
func UUID() string {
	uuid := make([]byte, 16)
	_, err := rand.Read(uuid)
	if err != nil {
		log.Panicln(err)
	}
	uuid[6] = (uuid[6] & 0x0f) | 0x40
	uuid[8] = (uuid[8] & 0x3f) | 0x80
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}
