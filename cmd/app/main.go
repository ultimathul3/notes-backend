package main

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/ultimathul3/notes-backend/internal/app"
	"github.com/ultimathul3/notes-backend/internal/config"
)

func init() {
	file, err := os.OpenFile("log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(io.MultiWriter(os.Stdout, file))
}

func main() {
	cfg, err := config.ReadEnvFile()
	if err != nil {
		log.Fatal(err)
	}

	app.Run(cfg)
}
