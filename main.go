package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

var introData StoryItem
var newyorkData StoryItem
var denverData StoryItem
var seanKellyData StoryItem
var markBatesData StoryItem
var debateData StoryItem
var homeData StoryItem

func main() {

	setData()

	mux := http.NewServeMux()
	tmpl = template.Must(template.ParseFiles("templates/index.gohtml"))

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/intro", intro)
	mux.HandleFunc("/denver", denver)
	mux.HandleFunc("/new-york", newYork)
	mux.HandleFunc("/home", home)
	mux.HandleFunc("/debate", debate)
	mux.HandleFunc("/sean-kelly", seanKelly)
	mux.HandleFunc("/mark-bates", markBates)

	log.Fatal(http.ListenAndServe(":8080", mux))

}

func intro(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, introData)
}

func denver(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, denverData)
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, homeData)
}

func debate(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, debateData)
}

func seanKelly(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, seanKellyData)
}

func markBates(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, markBatesData)
}

func newYork(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, newyorkData)
}

func setData() {
	data, err := parseJson("gopher.json")
	if err != nil {
		fmt.Println(err)
	}
	introData = data.Intro
	newyorkData = data.NewYork
	denverData = data.Denver
	seanKellyData = data.SeanKelly
	markBatesData = data.MarkBates
	debateData = data.Debate
	homeData = data.Home
}
