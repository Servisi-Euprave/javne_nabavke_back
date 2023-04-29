package repository

import "javne_nabavke_back/model"

type INabavkaRepo interface {
	InsertProcurement(nabavka *model.Procurement) error
	InsertProcurementPlan(nabavka *model.ProcurementPlan) error
	GetProcurements() ([]*model.Procurement, error)
}
