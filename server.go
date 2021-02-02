package main

import (
	"fmt"
	"net/http"

	router "github/kaji2002/http"

	"github/kaji2002/controller"
	"github/kaji2002/service"
	"github/kaji2002/repository"
)

var (
	postRepository repository.PostRepository = repository.NewFirestoreRepository()
	postService service.PostService          = service.NewPostService(postRepository)
	// controllerを定義 PostController型
	postController controller.PostController = controller.NewPostController(postService)
	// routerを定義 今回はchi Router型
	httpRouter router.Router                 = router.NewChiRouter()
)

func main() {
	// 新規ルーターを作成
	const port string = ":8080"
	// routeを定義
	// 『localhost:8080/』をやるとrequestが送られ、responseが画面に出力される
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "UP and runnigng・・・・")
	})
	// 投稿一覧表示機能
	httpRouter.GET("/posts", postController.GetPosts)
	// 投稿機能
	httpRouter.POST("/posts", postController.AddPost)
	// Serverを立ち上げる
	httpRouter.SERVE(port)
}
