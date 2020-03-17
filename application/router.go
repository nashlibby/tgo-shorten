/**
 * Created by nash.tang.
 * Author: nash.tang <112614251@qq.com>
 */

package application

import "github.com/gorilla/mux"

type Router struct {
	MuxRoute *mux.Router
	Handler  *Handler
}

func NewRouter() *Router{
	return &Router{
		MuxRoute: mux.NewRouter(),
	}
}

func (r *Router) Bind() {
	r.MuxRoute.HandleFunc("/api/v1/shorten", r.Handler.CreateShortLink).Methods("POST")
	r.MuxRoute.HandleFunc("/api/v1/info", r.Handler.GetShortLink).Methods("GET")
	r.MuxRoute.HandleFunc("/{short_link:[a-zA-Z0-9]{1,11}}", r.Handler.Redirect).Methods("GET")
}
