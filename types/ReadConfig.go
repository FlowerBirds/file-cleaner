package types

import (
	"log"
	"os"
)

type ReadConfig struct {
}

func (ReadConfig) Println() {
	log.Printf("Read config.")
}

func (ReadConfig) Getenv(key string) string {
	return os.Getenv(key)
}
