package repository

import "javne_nabavke_back/model"

type INabavkaRepo interface {
	AddNabavka(nabavka *model.Nabavka) error
}
