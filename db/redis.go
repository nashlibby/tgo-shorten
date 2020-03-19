/**
 * Created by nash.tang.
 * Author: nash.tang <112614251@qq.com>
 */

package db

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"tgo-shorten/util"
	"time"
)

const (
	UrlHashKey         = "url_hash:%s:url"
	ShortLinkDetailKey = "short_link:%s:detail"
)

type RedisCli struct {
	Cli *redis.Client
}

type UrlDetail struct {
	Url        string        `json:"url"`
	CreatedAt  string        `json:"created_at"`
	Expiration time.Duration `json:"expiration"`
}

func NewRedisCli() *RedisCli {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if _, err := client.Ping().Result(); err != nil {
		panic(err)
	}

	return &RedisCli{
		Cli: client,
	}
}

// 生成短地址
func (r *RedisCli) Shorten(url string, exp int64) (string, error) {
	hashArr, err := util.UrlEncrypt(url)
	if err != nil {
		return "", err
	}
	hashUrl := hashArr[0]
	res, err := r.Cli.Get(fmt.Sprintf(UrlHashKey, hashUrl)).Result()
	if err == redis.Nil {

	} else if err != nil {
		return "", err
	} else {
		if res != "{}" {
			return res, nil
		}
	}

	err = r.Cli.Set(fmt.Sprintf(UrlHashKey, hashUrl), url, time.Minute*time.Duration(exp)).Err()
	if err != nil {
		return "", err
	}

	detail, err := json.Marshal(
		&UrlDetail{
			Url:        url,
			CreatedAt:  time.Now().Format("2006-01-02 15:04:05"),
			Expiration: time.Duration(exp),
		})
	if err != nil {
		return "", err
	}

	err = r.Cli.Set(fmt.Sprintf(ShortLinkDetailKey, hashUrl), detail, time.Minute*time.Duration(exp)).Err()
	if err != nil {
		return "", err
	}

	return hashUrl, nil
}

// 短地址详情
func (r *RedisCli) ShortLinkInfo(hashUrl string) (interface{}, error) {
	res, err := r.Cli.Get(fmt.Sprintf(ShortLinkDetailKey, hashUrl)).Result()
	if err == redis.Nil {
		return nil, err
	} else if err != nil {
		return nil, err
	} else {
		var urlDetail UrlDetail
		err := json.Unmarshal([]byte(res), &urlDetail)
		if err != nil {
			return nil, err
		} else {
			return urlDetail, nil
		}
	}
}

// 获取原地址
func (r *RedisCli) UnShorten(hashUrl string) (string, error) {
	res, err := r.Cli.Get(fmt.Sprintf(UrlHashKey, hashUrl)).Result()
	if err == redis.Nil {
		return "", err
	} else if err != nil {
		return "", err
	} else {
		return res, nil
	}
}
