package main

import (
	"io"
	"log"
	"os"

	"github.com/AbishSowrirajan/finleap/models"
	"github.com/AbishSowrirajan/finleap/routers"
)

func init() {

	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalln("Failed to open log file :", err)
	}

	multi := io.MultiWriter(file, os.Stdout)

	_ = log.New(multi,
		"PREFIX: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {

	models.Init()
	routers.Run()
}
