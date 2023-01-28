package repository

import (
	"errors"
	"time"
)

func AddToken(token string, id string, hour int64) error {
	_, err := rdb[0].Set(token, id, time.Duration(hour)*time.Hour).Result()
	if err != nil {
		return err
	} else {
		return nil
	}
}

func GetId(token string) (string, error) {
	res, err := rdb[0].Get(token).Result()
	if err != nil {
		return "", err
	}
	return res, nil
}

func DelToken(token string) error {
	_, err := rdb[0].Get(token).Result()
	if err != nil {
		return errors.New("no data to delete")
	} else {
		rdb[0].Del(token)
		return nil
	}
}
