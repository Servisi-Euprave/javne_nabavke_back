package repository

import "javne_nabavke_back/model"

type INabavkaRepo interface {
	InsertProcurement(nabavka *model.Procurement) error
	InsertProcurementPlan(nabavka *model.ProcurementPlan) error
	GetProcurements() ([]*model.Procurement, error)
	GetProcurementPlans(companyPiB string) ([]*model.ProcurementPlan, error)
	GetCompanyProcurements(companyPiB string) ([]*model.Procurement, error)
	GetAllProcurements() ([]*model.Procurement, error)
	DeclareWinner(companyPib string, id string) error
}
