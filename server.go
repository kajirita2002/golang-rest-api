package main

import (
	"fmt"
	"net/http"

	router "github/kaji2002/http"

	"github/kaji2002/controller"
)

var (
	postController controller.PostController = controller.NewPostController()
	httpRouter     router.Router             = router.NewChiRouter()
)

func main() {
	// 新規ルーターを作成
	const port string = ":8080"
	// routeを定義
	// 『localhost:8080/』をやるとrequestが送られ、responseが画面に出力される
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "UP and runnigng・・・・")
	})
	//投稿一覧表示機能
	httpRouter.GET("/posts", postController.GetPosts)
	// 投稿機能
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(port)
}
