package main

import (
	"encoding/json"
	"fmt"
	"goblog/data"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("./templates"))
	mux.Handle("/static/", http.StripPrefix("/static", files))
	mux.HandleFunc("/", index)

	s := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	log.Fatal(s.ListenAndServe())
	art, err := data.GetArticleByID(1)
	if err != nil {
		log.Fatal(err)
	}
	artBytes, err := json.MarshalIndent(art, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(artBytes))
}
