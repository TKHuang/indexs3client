package main

import (
	"log"
	"os"

	"github.com/uc-cdis/indexs3client/handlers"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	s3objectURL := os.Getenv("INPUT_URL")
	log.Printf("INPUT_URL: %s", s3objectURL)
	err := handlers.IndexS3Object(s3objectURL)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
