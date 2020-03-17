/**
 * Created by nash.tang.
 * Author: nash.tang <112614251@qq.com>
 */

package application

import (
	"TgoShorten/db"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/validator.v2"
	"net/http"
)

type Handler struct {
}

type ShortenReq struct {
	Url        string `json:"url" validate:"nonzero"`
	Expiration int64  `json:"expiration" validate:"min=0"`
}

type ShortenResp struct {
	ShortLink string `json:"short_link"`
}

// 生成短链接
func (h *Handler) CreateShortLink(w http.ResponseWriter, r *http.Request) {
	var req ShortenReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		RespondWithError(w, &StatusError{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("parse parameters failed %v", r.Body),
		})
		return
	}
	if err := validator.Validate(req); err != nil {
		RespondWithError(w, &StatusError{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("parse parameters failed %v", req),
		})
		return
	}
	defer r.Body.Close()

	cli := db.NewRedisCli()
	res, err := cli.Shorten(req.Url, req.Expiration)
	if err != nil {
		RespondWithError(w, err)
	} else {
		RespondWithJson(w, http.StatusCreated, ShortenResp{ShortLink:res})
	}
}

// 获取短链接
func (h *Handler) GetShortLink(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query()
	url := param.Get("url")
	cli := db.NewRedisCli()
	res, err := cli.ShortLinkInfo(url)
	if err != nil {
		RespondWithError(w, err)
	} else {
		RespondWithJson(w, http.StatusOK, res)
	}
}

// 地址跳转
func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	cli := db.NewRedisCli()
	res, err := cli.UnShorten(param["url"])
	if err != nil {
		RespondWithError(w, err)
	} else {
		http.Redirect(w, r, res, http.StatusTemporaryRedirect)
	}
}
