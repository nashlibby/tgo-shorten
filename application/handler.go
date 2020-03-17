/**
 * Created by nash.tang.
 * Author: nash.tang <112614251@qq.com>
 */

package application

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
}

type ShortenReq struct {
	Url        string `json:"url"`
	Expiration int64  `json:"expiration"`
}

type ShortenResp struct {
	ShortLink string `json:"short_link"`
}

func (h *Handler) CreateShortLink(w http.ResponseWriter, r *http.Request) {
	var req ShortenReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return
	}
	defer r.Body.Close()

	fmt.Printf("%v\n", req)
}

func (h *Handler) GetShortLink(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query()
	s := param.Get("shortlink")

	fmt.Printf("%s\n", s)
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	fmt.Printf("%s\n", param["short_link"])
}
