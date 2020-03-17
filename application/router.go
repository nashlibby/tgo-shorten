/**
 * Created by nash.tang.
 * Author: nash.tang <112614251@qq.com>
 */

package application

import (
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

type Router struct {
	MuxRoute   *mux.Router
	Handler    *Handler
	Middleware *Middleware
}

// 构造函数
func NewRouter() *Router {
	return &Router{
		MuxRoute: mux.NewRouter(),
	}
}

// 绑定路由
func (r *Router) Bind() {
	m := alice.New(r.Middleware.LogHandler, r.Middleware.RecoverHandler)
	r.MuxRoute.Handle("/api/v1/shorten", m.ThenFunc(r.Handler.CreateShortLink)).Methods("POST")
	r.MuxRoute.Handle("/api/v1/info", m.ThenFunc(r.Handler.GetShortLink)).Methods("GET")
	r.MuxRoute.Handle("/{url:[a-zA-Z0-9]{1,11}}", m.ThenFunc(r.Handler.Redirect)).Methods("GET")
}
