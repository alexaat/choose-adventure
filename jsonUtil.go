package main

import (
	"encoding/json"
	"os"
)

func parseJson(file string) (Book, error) {
	var book Book = Book{}

	//Read File
	storiesBytes, error := os.ReadFile(file)
	if error != nil {
		return book, error
	}

	//Convert Json to Struct
	err := json.Unmarshal(storiesBytes, &book)
	if err != nil {
		return book, err
	}
	return book, nil
}

type Book struct {
	Intro     StoryItem `json:"intro"`
	NewYork   StoryItem `json:"new-york"`
	Debate    StoryItem `json:"debate"`
	SeanKelly StoryItem `json:"sean-kelly"`
	MarkBates StoryItem `json:"mark-bates"`
	Denver    StoryItem `json:"denver"`
	Home      StoryItem `json:"home"`
}

type StoryItem struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Link   `json:"options"`
}

type Link struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
