package main

import (
	"html/template"
	//"io/ioutil"
	"net/http"
)

func scorecardSearchHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/table.html")
    t.Execute(w, nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/index.html")
    t.Execute(w, nil)
}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/scorecard_search", scorecardSearchHandler) 
	http.ListenAndServe(":8000", nil)
}
