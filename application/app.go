/**
 * Created by nash.tang.
 * Author: nash.tang <112614251@qq.com>
 */

package application

import (
	"encoding/json"
	"log"
	"net/http"
)

type App struct {
}

// 构造函数
func NewApp() *App {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	return &App{
	}
}

func RespondWithError(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case Error:
		log.Printf("HTTP %d - %s", e.Status(), e)
		RespondWithJson(w, e.Status(), e.Error())
	default:
		RespondWithJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
}

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	resp, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(resp)
}

// 启动服务
func (a *App) Run(addr string) {
	router := NewRouter()
	router.Bind()
	log.Fatal(http.ListenAndServe(addr, router.MuxRoute))
}
