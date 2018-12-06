package main

import (
	"encoding/json"
	"fmt"
	"goblog/data"
	"log"
)

func main() {
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
