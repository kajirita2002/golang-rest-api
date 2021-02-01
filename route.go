package main

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var (
	posts []Post
)

func init() {
	posts = []Post{Post{Id: 1, Title: "Title 1", Text: "Text 1"}}
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result, error := json.Marshal(posts)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error: Error marshalling the post array"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
