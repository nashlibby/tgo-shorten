/**
 * Created by nash.tang.
 * Author: nash.tang <112614251@qq.com>
 */

package db

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gravityblast/go-base62"
	"log"
	"time"
)

const (
	UrlIdKey           = "next.url.id"
	ShortLinkKey       = "short_link:%s:url"
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

func (r *RedisCli) Shorten(url string, exp int64) (string, error) {
	hashVal := toSha1(url)
	res, err := r.Cli.Get(fmt.Sprintf(UrlHashKey, hashVal)).Result()
	if err == redis.Nil {

	} else if err != nil {
		return "", err
	} else {
		if res == "{}" {
		} else {
			return res, nil
		}
	}

	err = r.Cli.Incr(UrlIdKey).Err()
	if err != nil {
		return "", err
	}

	id, err := r.Cli.Get(UrlIdKey).Int()
	if err != nil {
		return "", err
	}

	eid := base62.Encode(id)
	log.Println(eid)

	err = r.Cli.Set(fmt.Sprintf(ShortLinkKey, eid), url, time.Minute*time.Duration(exp)).Err()
	if err != nil {
		return "", err
	}

	err = r.Cli.Set(fmt.Sprintf(UrlHashKey, hashVal), eid, time.Minute*time.Duration(exp)).Err()
	if err != nil {
		return "", err
	}

	detail, err := json.Marshal(
		&UrlDetail{
			Url:        url,
			CreatedAt:  time.Now().String(),
			Expiration: time.Duration(exp),
		})
	if err != nil {
		return "", err
	}

	err = r.Cli.Set(fmt.Sprintf(ShortLinkDetailKey, eid), detail, time.Minute*time.Duration(exp)).Err()
	if err != nil {
		return "", err
	}

	return eid, nil
}

func (r *RedisCli) ShortLinkInfo(eid string) (interface{}, error) {
	res, err := r.Cli.Get(fmt.Sprintf(ShortLinkDetailKey, eid)).Result()
	if err == redis.Nil {
		return "", err
	} else if err != nil {
		return "", err
	} else {
		return res, nil
	}
}

func (r *RedisCli) UnShorten(eid string) (string, error) {
	res, err := r.Cli.Get(fmt.Sprintf(ShortLinkKey, eid)).Result()
	if err == redis.Nil {
		return "", err
	} else if err != nil {
		return "", err
	} else {
		return res, nil
	}
}

func toSha1(data string) string {
	s := sha1.New()
	s.Write([]byte(data))
	return hex.EncodeToString(s.Sum([]byte("")))
}
