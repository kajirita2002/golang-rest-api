package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// muxRouter型
type muxRouter struct{}

var (
	muxDispatcher = mux.NewRouter()
)

// Router型(interfaceを持つ)のrouterを作成する
func NewMuxRouter() Router {
	// muxRouterのインスタンスを返す
	return &muxRouter{}
}
// method達
func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	// 実際の実装を書く
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

func (*muxRouter) SERVE(port string) {
	fmt.Println("Mux HTTP server runnning on port %v", port)
	http.ListenAndServe(port, muxDispatcher)
}
