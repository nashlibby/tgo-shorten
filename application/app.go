/**
 * Created by nash.tang.
 * Author: nash.tang <112614251@qq.com>
 */

package application

import (
	"log"
	"net/http"
)

type App struct {
}

func NewApp() *App {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	return &App{
	}
}

func (a *App) Run(addr string) {
	router := NewRouter()
	router.Bind()
	log.Fatal(http.ListenAndServe(addr, router.MuxRoute))
}
