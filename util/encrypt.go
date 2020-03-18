/**
 * Created by nash.tang.
 * Author: nash.tang <112614251@qq.com>
 */

package util

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"strconv"
)

const (
	VAL   = 0x3FFFFFFF
	INDEX = 0x0000003D
)

var (
	alphabet = []byte("abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

// url加密 返回4个短地址
func UrlEncrypt(url string) ([4]string, error) {
	md5Str := StrToMd5(url)
	var tempVal int64
	var result [4]string
	var tempUri []byte
	for i := 0; i < 4; i++ {
		tempSubStr := md5Str[i*8 : (i+1)*8]
		hexVal, err := strconv.ParseInt(tempSubStr, 16, 64)
		if err != nil {
			return result, nil
		}
		tempVal = int64(VAL) & hexVal
		var index int64
		tempUri = []byte{}
		for i := 0; i < 6; i++ {
			index = INDEX & tempVal
			tempUri = append(tempUri, alphabet[index])
			tempVal = tempVal >> 5
		}
		result[i] = string(tempUri)
	}
	return result, nil
}

// sha1加密
func StrToSha1(data string) string {
	s := sha1.New()
	s.Write([]byte(data))
	return hex.EncodeToString(s.Sum([]byte("")))
}

// md5加密
func StrToMd5(data string) string {
	m := md5.New()
	m.Write([]byte(data))
	c := m.Sum(nil)
	return hex.EncodeToString(c)
}
