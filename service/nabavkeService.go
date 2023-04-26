package service

import (
	"javne_nabavke_back/repository"
	"log"
)

type NabavkaService struct {
	l    *log.Logger
	repo repository.INabavkaRepo
}

func NewNabavkaService(l *log.Logger, repo repository.INabavkaRepo) *NabavkaService {
	return &NabavkaService{
		l:    l,
		repo: repo,
	}

}
