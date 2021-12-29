package types

import (
	"io"
	"log"
	"os"
)

type ReadConfig struct {
	LogDir string
}

func (ReadConfig) Println() {
	log.Printf("Read config.")
}

func (ReadConfig) Getenv(key string) string {
	return os.Getenv(key)
}

func (c *ReadConfig) Init(dir string) {
	logDir := dir + "logs"
	_, err := os.Stat(logDir)
	if err != nil {
		os.Mkdir(logDir, os.ModePerm)
		log.Printf("Create logs file: " + logDir)
	}
	c.LogDir = logDir
	logfile, err := os.OpenFile(logDir+"/file-cleaner.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		log.Fatalf("create log file error" + err.Error())
	}
	multiWriter := io.MultiWriter(os.Stdout, logfile)
	log.SetOutput(multiWriter)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
}
