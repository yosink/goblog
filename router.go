package main

import (
	"fmt"
	"goblog/data"
	"html/template"
	"net/http"
)

var templateDir = "./templates/"

func index(w http.ResponseWriter, r *http.Request) {
	arts, err := data.GetArticleList()
	//fmt.Printf("%#v", arts)
	////os.Exit(1)
	if err != nil {
		fmt.Fprint(w, "get article list error")
	}
	t := template.Must(template.ParseFiles(templateDir + "index.html"))
	t.Execute(w, &arts)
}
