package main

import (
	"io"
	"log"
	"os"

	"github.com/finleap/models"
)

func init() {

	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalln("Failed to open log file", output, ":", err)
	}

	multi := io.MultiWriter(file, os.Stdout)

	logger := log.New(multi,
		"PREFIX: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {

	models.Init()
}
