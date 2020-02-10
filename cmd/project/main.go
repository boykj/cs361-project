package main

import (
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

func main() {

	fileServer := http.FileServer(http.Dir("./internal/ui/static/"))
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)

	godotenv.Load()
	PORT := os.Getenv("PORT")
	log.Println("Starting a server on :" + PORT)
	err := http.ListenAndServe(":"+PORT, mux)
	log.Fatal(err)
}