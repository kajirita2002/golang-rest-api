package controller

import (
	"encoding/json"
	"github/kaji2002/errors"
	"net/http"

	"github/kaji2002/entity"
	"github/kaji2002/service"
)

type controller struct{}

// 定義が必要
var (
	postService service.PostService = service.NewPostService()
)

type PostController interface {
	GetPosts(w http.ResponseWriter, r *http.Request)
	AddPost(w http.ResponseWriter, r *http.Request)
}

func NewPostController() PostController {
	return &controller{}
}

// 投稿一覧機能
func (*controller) GetPosts(w http.ResponseWriter, r *http.Request) {
	// responseのcontent-typeを設定
	w.Header().Set("Content-Type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error getting the post"})
	}

	// errなしの場合は200をかえし
	w.WriteHeader(http.StatusOK)
	// responseでresultを返す
	json.NewEncoder(w).Encode(posts)
}

// 投稿機能
func (*controller) AddPost(w http.ResponseWriter, r *http.Request) {
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
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error unmarshalling the request"})
		return
	}
	err1 := postService.Validate(&post)
	if err1 != nil {
		// 500エラーを出す
		w.WriteHeader(http.StatusInternalServerError)
		// エラーメッセージ
		json.NewEncoder(w).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}

	result, err2 := postService.Create(&post)
	if err2 != nil {
		// 500エラーを出す
		w.WriteHeader(http.StatusInternalServerError)
		// エラーメッセージ
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error saving the post"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
