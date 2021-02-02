// controllerは実際の実装ではなくエラー処理を行う
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
	postService service.PostService
)

// interface
type PostController interface {
	// これを実装しないといけない
	GetPosts(w http.ResponseWriter, r *http.Request)
	AddPost(w http.ResponseWriter, r *http.Request)
}

// PostControllerのinterfaceを搭載した
func NewPostController(service service.PostService) PostController {
	// serviceが変わっても対応可能
	postService = service
	// 抽象化したserviceを搭載したcontrollerを返す
	return &controller{}
}

// PostControllerの実際の実装
// 投稿一覧機能
func (*controller) GetPosts(w http.ResponseWriter, r *http.Request) {
	// responseのcontent-typeを設定
	w.Header().Set("Content-Type", "application/json")
	// 全検索(抽象化)
	posts, err := postService.FindAll()
	// エラー処理の実装
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
