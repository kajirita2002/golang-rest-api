package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type chiRouter struct{}

var (
	chiDispathcer = chi.NewRouter()
)

func NewChiRouter() Router {
	return &chiRouter{}
}

func (*chiRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispathcer.Get(uri, f)
}
func (*chiRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispathcer.Post(uri, f)
}

func (*chiRouter) SERVE(port string) {
	fmt.Println("Chi HTTP server runnning on port %v", port)
	http.ListenAndServe(port, chiDispathcer)
}
