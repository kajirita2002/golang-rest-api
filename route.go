package main

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"./entity"
	"./repository"
)

// 定義が必要
var (
	repo repository.PostRepository = repository.NewPostRepository()
)

// 投稿一覧機能
func getPosts(w http.ResponseWriter, r *http.Request) {
	// responseのcontent-typeを設定
	w.Header().Set("Content-Type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`error: "Error getting the posts"`))
	}

	// errなしの場合は200をかえし
	w.WriteHeader(http.StatusOK)
	// responseでresultを返す
	json.NewEncoder(w).Encode(posts)
}

// 投稿機能
func addPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// postの箱を準備
	var post entity.Post
	// requestを読み込んでpostに格納する(validate)
	err := json.NewDecoder(r.Body).Decode(&post)
	// errがあれば
	if err != nil {
		// 500エラーを出す
		w.WriteHeader(http.StatusInternalServerError)
		// エラーメッセージ
		w.Write([]byte(`{"error: "Error unmarshalling the request"}`))
		return
	}
	post.ID = rand.Int63()
	// postsをpostを追加して更新
	repo.Save(&post)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}
