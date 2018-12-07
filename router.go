package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Println(r.UserAgent())
	_, err := fmt.Fprint(w, "blog index")
	if err != nil {
		log.Fatal(err)
	}
}
