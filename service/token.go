package service

import (
	"crypto/md5"
	"fmt"
	"strconv"

	"github.com/Iscolito/Vshare/repository"
)

const validtime = 24

func Tokenize(rawStr string) string {
	data := []byte(rawStr)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func SendToken(token string, id int64) error {
	return repository.AddToken(token, strconv.FormatInt(id, 10), validtime)
}

func GetTokenId(token string) (int64, error) {
	strid, err := repository.GetId(token)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(strid, 10, 64)
}
