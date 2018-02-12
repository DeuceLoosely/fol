package main

import (
	"log"
	"net/http"
)

func main() {
	FileServe := http.FileServer(http.Dir("static"))
	ImageServe := http.FileServer(http.Dir("static/images"))
	http.Handle("/", FileServe)
	http.Handle("/images", ImageServe)

	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
