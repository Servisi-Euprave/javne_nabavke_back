package repository

import "javne_nabavke_back/model"

type INabavkaRepo interface {
	AddProcurement(nabavka *model.Procurement) error
	GetProcurements() ([]model.Procurement, error)
}
