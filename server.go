package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// 新規ルーターを作成
	router := mux.NewRouter()
	const port string = ":8080"
	// routeを定義
	// 『localhost:8080/』をやるとrequestが送られ、responseが画面に出力される
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "UP and runnigng・・・・")
	})
	router.HandleFunc("/posts", getPosts).Methods("GET")
	// logとポート番号を出力する
	log.Println("Server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}