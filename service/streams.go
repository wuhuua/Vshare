package service

import (
	"errors"

	"github.com/Iscolito/Vshare/repository"
	"github.com/Iscolito/Vshare/util"
)

func GetStreams() ([]repository.Video, error) {
	VideoDao := repository.NewVideoDaoInstance()
	streams, err := VideoDao.GetVideos(30)
	if err == nil {
		if streams == nil {
			util.Logger.Error("no videos")
			return nil, errors.New("no videos")
		} else {
			return streams, nil
		}
	} else {
		util.Logger.Error("find videos process error")
		return nil, errors.New("find videos process error")
	}
}
