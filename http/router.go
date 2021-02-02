// packageはrouterにする (interfaceのため)
package router

import "net/http"

// interfaceを定義する routerは必ず以下を実装しなければならない
type Router interface {
	// 引数にuriとfuncを代入
	GET(uri string, f func(w http.ResponseWriter, r *http.Request))
	POST(uri string, f func(w http.ResponseWriter, r *http.Request))
	SERVE(port string)
}
