package main

import (
	"log"
	"miniProject_golang/hendler"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", hendler.HomeHendler)
	mux.HandleFunc("/createTask", hendler.CreateTaskHendler)
	mux.HandleFunc("/product", hendler.ProductHendler)
	mux.HandleFunc("/post-get", hendler.PostGet)

	assetsfileServer := http.FileServer(http.Dir("assets"))
	datafileServer := http.FileServer(http.Dir("data"))
	mux.Handle("/assets/", http.StripPrefix("/assets", assetsfileServer))
	mux.Handle("/data/", http.StripPrefix("/data", datafileServer))

	log.Println("Starting web on port 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
